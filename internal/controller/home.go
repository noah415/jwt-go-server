package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noah415/jwt-go-server/internal/application"
)

func GetHome(c *gin.Context) {
	resp := application.GetHome()

	c.JSON(200, resp)
}
