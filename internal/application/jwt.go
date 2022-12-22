package application

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	RefreshType int16 = iota
	ReadType
	WriteType
)

func ParseTokenType(typ int16) (string, error) {
	switch typ {
	case RefreshType:
		return "refresh", nil
	case ReadType:
		return "read", nil
	case WriteType:
		return "write", nil
	default:
		return "", errors.New("unknown token type given to parser")
	}
}

func getTokenLifeSpans() (int, int, error) {
	var refresh int
	var access int
	var refreshStr string
	var accessStr string

	var refreshConvErr error
	var accessConvErr error

	if refreshStr = os.Getenv("REFRESH_TOKEN_LIFE"); refreshStr == "" {
		return -1, -1, errors.New("could not find \"REFRESH_TOKEN_LIFE\"")
	}

	if refresh, refreshConvErr = strconv.Atoi(refreshStr); refreshConvErr != nil {
		return -1, -1, errors.New("error converting the refresh string to an int: " + refreshConvErr.Error())
	}

	if accessStr = os.Getenv("ACCESS_TOKEN_LIFE"); accessStr == "" {
		return -1, -1, errors.New("could not find \"ACCESS_TOKEN_LIFE\"")
	}

	if access, accessConvErr = strconv.Atoi(accessStr); accessConvErr != nil {
		return -1, -1, errors.New("error converting the access string to an int: " + accessConvErr.Error())
	}

	return refresh, access, nil
}

func createJwtClaims(username string, typ string) (jwt.MapClaims, error) {
	var refreshLifeSpan int
	var accessLifeSpan int
	var getTokenErr error

	if refreshLifeSpan, accessLifeSpan, getTokenErr = getTokenLifeSpans(); getTokenErr != nil {
		return jwt.MapClaims{}, errors.New("error retrieving token life spans: " + getTokenErr.Error())
	}

	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["typ"] = typ

	if typ == "refresh" {
		claims["exp"] = time.Now().Add(time.Minute * time.Duration(refreshLifeSpan)).Unix()
	} else {
		claims["exp"] = time.Now().Add(time.Minute * time.Duration(accessLifeSpan)).Unix()
	}

	return claims, nil
}

func CreateToken(username string, typEnum int16) (string, error) {
	var bigBang string
	var rotationStr string
	var rotation int

	var signedString string
	var signingErr error

	var typ string
	var typErr error

	var claims jwt.MapClaims
	var claimsErr error

	if typ, typErr = ParseTokenType(typEnum); typErr != nil {
		return "", errors.New("error parsing the token type: " + typErr.Error())
	}

	if claims, claimsErr = createJwtClaims(username, typ); claimsErr != nil {
		return "", errors.New("error setting claims: " + claimsErr.Error())
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if bigBang = os.Getenv("TOKEN_SECRET"); bigBang == "" {
		return "", errors.New("could not find \"TOKEN_SECRET\" from env file")
	}

	if rotationStr = os.Getenv("SECRET_ROTATION"); rotationStr == "" {
		return "", errors.New("could not find \"SECRET_ROTATION\" from env file")
	}

	rotation, _ = strconv.Atoi(rotationStr)
	shift := time.Now().Unix() / int64(rotation) * int64(rotation)
	secret := bigBang + strconv.Itoa(int(shift))

	if signedString, signingErr = token.SignedString([]byte(secret)); signingErr != nil {
		return "", errors.New("could not sign the generated token with the given env secret: " + signingErr.Error())
	}

	return signedString, nil
}

func jwtParser(token *jwt.Token) (interface{}, error) {
	var rotationStr string
	var rotation int

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("token's signature has been tampered with")
	}

	if bigBang := os.Getenv("TOKEN_SECRET"); bigBang == "" {
		return nil, errors.New("failed to retrieve \"TOKEN_SECRET\" from the env file during authorization")
	} else {
		if rotationStr = os.Getenv("SECRET_ROTATION"); rotationStr == "" {
			return "", errors.New("could not find \"SECRET_ROTATION\" from env file")
		}

		rotation, _ = strconv.Atoi(rotationStr)
		shift := time.Now().Unix() / int64(rotation) * int64(rotation)
		secret := bigBang + strconv.Itoa(int(shift))
		return []byte(secret), nil
	}
}

func getClaim(claimStr string, claims jwt.MapClaims) (string, error) {
	claim := claims[claimStr].(string)

	if claim == "" {
		return "", errors.New("token does not contain the claim \"" + claimStr + "\"")
	}

	return claim, nil
}

func ValidateTokenAndRetrieveUsername(tokenString string) (string, error) {
	var usr string
	var usrErr error

	token, err := jwt.Parse(tokenString, jwtParser)

	if err != nil {
		return "", errors.New("error while parsing jwt: " + err.Error())
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if usr, usrErr = getClaim("username", claims); usrErr != nil {
			return "", usrErr
		}

		if !CheckUserExist(usr) {
			return "", errors.New("token contains invalid username")
		}

		return claims["username"].(string), nil
	} else {
		return "", errors.New("error while retrieving claims from jwt token")
	}
}
