package auth

import (
	"github.com/Osas997/go-portfolio/internal/domains/auth/controller"
	"github.com/Osas997/go-portfolio/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, controller controller.AuthController) {
	routes := router.Group("/auth")
	{
		routes.POST("/login", controller.Login)

		routes.GET("/me", middleware.AuthMiddleware(), func(ctx *gin.Context) {
			payload, _ := ctx.Get("user")
			ctx.JSON(200, gin.H{"data": payload})
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
