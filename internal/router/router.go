package router

import (
	"github.com/gin-gonic/gin"
	"github.com/noah415/jwt-go-server/internal/middleware"
)

var router = gin.Default()

func InitRouter() {
	setupMiddleware()
	getRoutes()
	router.SetTrustedProxies(nil)
	router.Run()
}

func getRoutes() {
	// Non-Authorized Routes
	InitHomeRoutes(router.Group("/"))
	InitRegisterRoutes(router.Group("/register"))
	InitLoginRoutes(router.Group("/login"))

	// Authorized Routes
	InitAuthorizeRoutes(router.Group("/auth", middleware.AuthorizeHandler()))
}

func setupMiddleware() {
	router.Use(gin.CustomRecovery(middleware.CustomRecoveryMiddleware()))
	router.Use(middleware.ErrorHandler())
}
