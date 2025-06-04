package controller

import (
	"github.com/Osas997/go-portfolio/internal/domains/projects/params"
	"github.com/Osas997/go-portfolio/internal/domains/projects/service"
	"github.com/Osas997/go-portfolio/internal/pkg/errorhandler"
	"github.com/Osas997/go-portfolio/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProjectControllerImpl struct {
	ProjectService service.ProjectService
	validate       *validator.Validate
}

func NewProjectController(projectService service.ProjectService, validate *validator.Validate) ProjectController {
	return &ProjectControllerImpl{ProjectService: projectService, validate: validate}
}

// Create implements ProjectController.
func (p *ProjectControllerImpl) Create(ctx *gin.Context) {
	var projectReq params.ProjectRequest

	if err := ctx.ShouldBind(&projectReq); err != nil {
		errorhandler.HandleError(ctx, errorhandler.NewBadRequestError("Invalid form data", err.Error()))
		return
	}

	if err := p.validate.Struct(projectReq); err != nil {
		errorhandler.HandleError(ctx, err)
		return
	}

	project, err := p.ProjectService.Create(&projectReq, ctx)
	if err != nil {
		errorhandler.HandleError(ctx, err)
		return
	}

	webResponse := utils.NewWebResponse("Project created successfully", project.ID)

	ctx.JSON(201, webResponse)
}

// Delete implements ProjectController
func (p *ProjectControllerImpl) Delete(ctx *gin.Context) {
	projectId := ctx.Param("id")
	err := p.ProjectService.Delete(projectId)

	if err != nil {
		errorhandler.HandleError(ctx, err)
		return
	}

	webResponse := utils.NewWebResponse("Project deleted successfully", nil)

	ctx.JSON(200, webResponse)
}

// FindAll implements ProjectController.
func (p *ProjectControllerImpl) FindAll(ctx *gin.Context) {
	projects, err := p.ProjectService.FindAll()
	if err != nil {
		errorhandler.HandleError(ctx, err)
		return
	}

	webResponse := utils.NewWebResponse("Projects found successfully", projects)

	ctx.JSON(200, webResponse)
}

// FindById implements ProjectController.
func (p *ProjectControllerImpl) FindById(ctx *gin.Context) {
	projectId := ctx.Param("id")
	project, err := p.ProjectService.FindById(projectId)
	if err != nil {
		errorhandler.HandleError(ctx, err)
		return
	}

	webResponse := utils.NewWebResponse("Project found successfully", project)

	ctx.JSON(200, webResponse)
}

// Update implements ProjectController.
func (p *ProjectControllerImpl) Update(ctx *gin.Context) {
	var projectReq params.UpdateProjectRequest

	if err := ctx.ShouldBind(&projectReq); err != nil {
		errorhandler.HandleError(ctx, errorhandler.NewBadRequestError("Invalid form data", err.Error()))
		return
	}

	if err := p.validate.Struct(projectReq); err != nil {
		errorhandler.HandleError(ctx, err)
		return
	}

	projectId := ctx.Param("id")
	updatedProject, err := p.ProjectService.Update(projectId, &projectReq, ctx)
	if err != nil {
		errorhandler.HandleError(ctx, err)
		return
	}

	webResponse := utils.NewWebResponse("Project updated successfully", updatedProject.ID)

	ctx.JSON(200, webResponse)
}
