package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noah415/Recibase-business-logic/internal/application"
)

func GetHome(c *gin.Context) {
	resp := application.GetHome()

	c.JSON(200, resp)
}
