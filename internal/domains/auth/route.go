package auth

import (
	"github.com/Osas997/go-portfolio/internal/domains/auth/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, controller controller.AuthController) {
	routes := router.Group("/auth")
	{
		routes.POST("/login", controller.Login)

		routes.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "Hello World!"})
		})
		// routes.POST("/register", handler.Register)

		// Protected routes
		// authRoutes := routes.Group("/")
		// authRoutes.Use(middleware.AuthMiddleware())
		// {
		// 	authRoutes.GET("/profile", handler.GetProfile)
		// }
	}
}
