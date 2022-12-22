package router

import (
	"github.com/gin-gonic/gin"
	"github.com/noah415/Recibase-business-logic/internal/middleware"
)

var router = gin.Default()

func InitRouter() {
	setupMiddleware()
	getRoutes()
	router.SetTrustedProxies(nil)
	router.Run()
}

func getRoutes() {
	InitHomeRoutes(router.Group("/"))
	InitRegisterRoutes(router.Group("/register"))
	InitLoginRoutes(router.Group("/login"))
	InitAuthorizeRoutes(router.Group("/auth", middleware.AuthorizeHandler()))
}

func setupMiddleware() {
	router.Use(gin.CustomRecovery(middleware.CustomRecoveryMiddleware()))
	router.Use(middleware.ErrorHandler())
}
