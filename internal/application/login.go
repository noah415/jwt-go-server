package application

import "errors"

type PostLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PostLoginResponse struct {
	Token string `json:"token"`
}

func validateCredentials(req *PostLoginRequest) bool {
	for _, value := range Users {
		if value.Username == req.Username &&
			value.Password == req.Password {
			return true
		}
	}
	return false
}

func PostLogin(req *PostLoginRequest) (PostLoginResponse, error) {
	var token string
	var tokenCreationErr error

	if valid := validateCredentials(req); !valid {
		return PostLoginResponse{}, errors.New("invalid credentials")
	}

	if token, tokenCreationErr = CreateToken(req.Username, RefreshType); tokenCreationErr != nil {
		return PostLoginResponse{}, tokenCreationErr
	}

	return PostLoginResponse{
		Token: token,
	}, nil
}
