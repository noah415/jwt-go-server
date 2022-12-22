package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors[0]

			message, code := parseError(err.Err)

			c.JSON(code, gin.H{
				"message": message,
			})
		}
	}
}

func parseError(err error) (string, int) {
	errType := strings.Split(err.Error(), ":")[0]
	errMessage := strings.Split(err.Error(), ":")[1]
	code := 0

	switch errType {
	case "Validation":
		code = 422
	case "Bad Request":
		code = 400
	case "Authorization":
		code = 401
	default:
		code = 500
	}

	return errMessage, code
}

func CustomRecoveryMiddleware() func(c *gin.Context, recovered interface{}) {
	return func(c *gin.Context, recovered interface{}) {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal Server Error",
		})
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
