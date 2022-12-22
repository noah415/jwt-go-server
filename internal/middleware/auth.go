package middleware

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/noah415/jwt-go-server/internal/application"
	"github.com/noah415/jwt-go-server/internal/exception"
)

func AuthorizeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var header string
		var tokenList []string

		if header = c.GetHeader("Authorization"); header == "" {
			c.Errors = append(c.Errors, c.Error(&exception.Exception{
				ErrType: exception.AuthorizationError,
				Err:     errors.New("no authentication token found in the header"),
			}))
			return
		}

		if tokenList = strings.Split(header, " "); len(tokenList) == 2 {
			username, err := application.ValidateTokenAndRetrieveUsername(tokenList[1])
			if err != nil {
				c.Errors = append(c.Errors, c.Error(&exception.Exception{
					ErrType: exception.AuthorizationError,
					Err:     errors.New("error during token validation: " + err.Error()),
				}))
				return
			}

			c.Set("username", username)
		} else {
			c.Errors = append(c.Errors, c.Error(&exception.Exception{
				ErrType: exception.AuthorizationError,
				Err:     errors.New("authorization header could not be parsed into two parts (bearer and token)"),
			}))
			return
		}

		c.Next()
	}
}
