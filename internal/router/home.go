package router

import (
	"github.com/gin-gonic/gin"
	"github.com/noah415/Recibase-business-logic/internal/controller"
)

func InitHomeRoutes(r *gin.RouterGroup) {
	r.GET("", controller.GetHome)
}
