package application

import "errors"

type PostRegisterRequest struct {
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	ConfPassword string `json:"confPassword" binding:"required"`
}

type PostRegisterResponse struct {
	Token string `json:"token"`
}

type User struct {
	Username string
	Password string
}

var Users = []User{
	{
		Username: "notsuka",
		Password: "password",
	},
}

func CheckUserExist(username string) bool {
	for _, value := range Users {
		if value.Username == username {
			return true
		}
	}
	return false
}

func validatePasswords(password string, confPassword string) bool {
	return password == confPassword
}

func PostRegister(req *PostRegisterRequest) (PostRegisterResponse, error) {
	var newUser User

	var token string
	var tokenCreationErr error

	if userExists := CheckUserExist(req.Username); userExists {
		return PostRegisterResponse{}, errors.New("username already exists")
	} else if validPassword := validatePasswords(req.Password, req.ConfPassword); !validPassword {
		return PostRegisterResponse{}, errors.New("passwords do not match")
	} else {
		newUser = User{
			Username: req.Username,
			Password: req.Password,
		}

		Users = append(Users, newUser)

		if token, tokenCreationErr = CreateToken(newUser.Username, RefreshType); tokenCreationErr != nil {
			return PostRegisterResponse{}, tokenCreationErr
		}

		return PostRegisterResponse{
			Token: token,
		}, nil
	}
}
