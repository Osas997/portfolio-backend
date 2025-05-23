package service

import (
	"github.com/Osas997/go-portfolio/internal/domains/projects/entity"
	"github.com/Osas997/go-portfolio/internal/domains/projects/params"
	"github.com/gin-gonic/gin"
)

type ProjectService interface {
	FindAll() ([]*params.ProjectResponse, error)
	FindById(projectId string) (*params.ProjectDetailResponse, error)
	Create(projectReq *params.ProjectRequest, ctx *gin.Context) (*entity.Projects, error)
	Update(projectId string, projectReq *params.UpdateProjectRequest, ctx *gin.Context) (*entity.Projects, error)
	Delete(projectId string) error
}
