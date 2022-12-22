package router

import (
	"github.com/gin-gonic/gin"
	"github.com/noah415/jwt-go-server/internal/controller"
)

func InitHomeRoutes(r *gin.RouterGroup) {
	r.GET("", controller.GetHome)
}
