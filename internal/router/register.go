package router

import (
	"github.com/gin-gonic/gin"
	"github.com/noah415/jwt-go-server/internal/controller"
)

func InitRegisterRoutes(r *gin.RouterGroup) {
	r.POST("", controller.PostRegister)
}
