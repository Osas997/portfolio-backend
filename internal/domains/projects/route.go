package auth

import (
	"github.com/Osas997/go-portfolio/internal/domains/projects/controller"
	"github.com/Osas997/go-portfolio/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, controller controller.ProjectController) {
	routes := router.Group("/project")
	router.Use(middleware.AuthMiddleware())
	{
		routes.GET("/", controller.FindAll)
	}
}
