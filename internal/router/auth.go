package router

import "github.com/gin-gonic/gin"

func InitAuthorizeRoutes(r *gin.RouterGroup) {
	// This is just an empty route that I was using to test authorized routes
	r.GET("")
}
