package controller

import (
	"github.com/Osas997/go-portfolio/internal/domains/projects/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProjectControllerImpl struct {
	ProjectService service.ProjectService
	validate       *validator.Validate
}

func NewAuthController(projectService service.ProjectService, validate *validator.Validate) ProjectController {
	return &ProjectControllerImpl{ProjectService: projectService, validate: validate}
}

// Create implements ProjectController.
func (p *ProjectControllerImpl) Create(ctx *gin.Context) {
	panic("unimplemented")
}

// Delete implements ProjectController.
func (p *ProjectControllerImpl) Delete(ctx *gin.Context) {
	panic("unimplemented")
}

// FindAll implements ProjectController.
func (p *ProjectControllerImpl) FindAll(ctx *gin.Context) {
	panic("unimplemented")
}

// FindById implements ProjectController.
func (p *ProjectControllerImpl) FindById(ctx *gin.Context) {
	panic("unimplemented")
}

// Update implements ProjectController.
func (p *ProjectControllerImpl) Update(ctx *gin.Context) {
	panic("unimplemented")
}
