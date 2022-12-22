package router

import (
	"github.com/gin-gonic/gin"
	"github.com/noah415/Recibase-business-logic/internal/controller"
)

func InitLoginRoutes(r *gin.RouterGroup) {
	r.POST("", controller.PostLogin)
}
