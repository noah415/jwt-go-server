package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/noah415/Recibase-business-logic/internal/application"
	"github.com/noah415/Recibase-business-logic/internal/exception"
)

func PostLogin(c *gin.Context) {
	var req application.PostLoginRequest
	var resp application.PostLoginResponse
	var err error

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Errors = append(c.Errors, c.Error(&exception.Exception{
			ErrType: exception.BadRequestError,
			Err:     errors.New("invalid request body"),
		}))
		return
	}

	if resp, err = application.PostLogin(&req); err != nil {
		c.Errors = append(c.Errors, c.Error(&exception.Exception{
			ErrType: exception.ValidationError,
			Err:     err,
		}))
		return
	}

	c.JSON(200, resp)
}
