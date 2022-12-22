package router

import "github.com/gin-gonic/gin"

func InitAuthorizeRoutes(r *gin.RouterGroup) {
	r.GET("")
	r.POST("/recipe")
}
