package projects

import (
	"github.com/Osas997/go-portfolio/internal/domains/projects/controller"
	"github.com/Osas997/go-portfolio/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, controller controller.ProjectController) {
	routesPublic := router.Group("/projects")
	routesPublic.GET("", controller.FindAll)

	routes := router.Group("/cms/projects")
	routes.Use(middleware.AuthMiddleware())
	{
		routes.GET("", controller.FindAll)
		routes.POST("", controller.Create)
		routes.GET(":id", controller.FindById)
		routes.PATCH(":id", controller.Update)
		routes.DELETE(":id", controller.Delete)
	}
}
