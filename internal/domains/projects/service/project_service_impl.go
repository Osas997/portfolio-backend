package service

import (
	"os"

	"github.com/Osas997/go-portfolio/internal/domains/projects/entity"
	"github.com/Osas997/go-portfolio/internal/domains/projects/params"
	"github.com/Osas997/go-portfolio/internal/domains/projects/repository"
	"github.com/Osas997/go-portfolio/internal/pkg/errorhandler"
	"github.com/Osas997/go-portfolio/internal/pkg/uploadfile"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProjectServiceImpl struct {
	repository              repository.ProjectRepository
	projectImagesRepository repository.ProjectImagesRepository
}

func NewProjectService(repository repository.ProjectRepository, projectImagesRepository repository.ProjectImagesRepository) ProjectService {
	return &ProjectServiceImpl{
		repository,
		projectImagesRepository,
	}
}

// Create implements ProjectService.
func (p *ProjectServiceImpl) Create(projectReq *params.ProjectRequest, ctx *gin.Context) (*entity.Projects, error) {
	project := entity.Projects{
		Title:   projectReq.Title,
		Content: projectReq.Content,
	}

	dst := "./uploads/thumbnails/" + uploadfile.FormatUpload(projectReq.Thumbnail)

	if err := ctx.SaveUploadedFile(projectReq.Thumbnail, dst); err != nil {
		return &entity.Projects{}, errorhandler.NewBadRequestError("gagal menyimpan file", err.Error())
	}

	project.Thumbnail = dst

	newProject, err := p.repository.Save(&project)
	if err != nil {
		return &entity.Projects{}, err
	}

	// Save All Images
	for _, image := range projectReq.Images {
		dst := "./uploads/images/" + uploadfile.FormatUpload(image)

		if err := ctx.SaveUploadedFile(image, dst); err != nil {
			return &entity.Projects{}, errorhandler.NewBadRequestError("gagal menyimpan file", err.Error())
		}

		projectImages := entity.ProjectImages{
			ProjectID: newProject.ID,
			Image:     dst,
		}

		_, err := p.projectImagesRepository.Save(&projectImages)
		if err != nil {
			return &entity.Projects{}, err
		}
	}

	return newProject, nil
}

// Delete implements ProjectService.
func (p *ProjectServiceImpl) Delete(userId string) error {
	project, err := p.repository.FindById(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errorhandler.NewNotFoundError("Project not found")
		}
		return err
	}

	// Hapus Image Lama
	if err := os.Remove(project.Thumbnail); err != nil && !os.IsNotExist(err) {
		return err
	}

	if project.ProjectImages != nil {
		for _, image := range project.ProjectImages {
			if err := os.Remove(image.Image); err != nil && !os.IsNotExist(err) {
				return err
			}
		}
	}

	if err := p.repository.Delete(project.ID.String()); err != nil {
		return err
	}

	return nil
}

// FindAll implements ProjectService.
func (p *ProjectServiceImpl) FindAll() ([]*params.ProjectResponse, error) {
	projects, err := p.repository.FindAll()
	if err != nil {
		return []*params.ProjectResponse{}, err
	}

	projectResponse := params.NewProjectResponse(projects)

	return projectResponse, nil
}

// FindById implements ProjectService.
func (p *ProjectServiceImpl) FindById(userId string) (*params.ProjectDetailResponse, error) {
	project, err := p.repository.FindById(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &params.ProjectDetailResponse{}, errorhandler.NewNotFoundError("Project not found")
		}
		return &params.ProjectDetailResponse{}, err
	}

	projectResponse := params.NewProjectDetailResponse(project)

	return projectResponse, nil
}

// Update implements ProjectService.
func (p *ProjectServiceImpl) Update(projectId string, projectReq *params.UpdateProjectRequest, ctx *gin.Context) (*entity.Projects, error) {
	project, err := p.repository.FindById(projectId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &entity.Projects{}, errorhandler.NewNotFoundError("Project not found")
		}
		return &entity.Projects{}, err
	}

	if projectReq.Thumbnail != nil {
		// Hapus Image Lama
		err := os.Remove(project.Thumbnail)
		if err != nil && !os.IsNotExist(err) {
			return &entity.Projects{}, err
		}

		dst := "./uploads/thumbnails/" + uploadfile.FormatUpload(projectReq.Thumbnail)
		if err := ctx.SaveUploadedFile(projectReq.Thumbnail, dst); err != nil {
			return &entity.Projects{}, errorhandler.NewBadRequestError("gagal menyimpan file", err.Error())
		}

		project.Thumbnail = dst
	}

	if projectReq.Images != nil {
		// Hapus Image Lama
		for _, oldImage := range project.ProjectImages {
			err := os.Remove(oldImage.Image)
			if err != nil && !os.IsNotExist(err) {
				return &entity.Projects{}, err
			}
		}

		if err := p.projectImagesRepository.DeleteAllByProjectId(projectId); err != nil {
			return &entity.Projects{}, err
		}

		projectImages := []entity.ProjectImages{}
		for _, image := range projectReq.Images {
			dst := "./uploads/images/" + uploadfile.FormatUpload(image)
			if err := ctx.SaveUploadedFile(image, dst); err != nil {
				return &entity.Projects{}, errorhandler.NewBadRequestError("gagal menyimpan file", err.Error())
			}

			projectImage := entity.ProjectImages{
				ProjectID: project.ID,
				Image:     dst,
			}
			projectImages = append(projectImages, projectImage)
		}
		project.ProjectImages = projectImages
	}

	if projectReq.Title != nil {
		project.Title = *projectReq.Title
	}

	if projectReq.Content != nil {
		project.Content = *projectReq.Content
	}

	updatedProject, err := p.repository.Save(project)
	if err != nil {
		return &entity.Projects{}, err
	}

	return updatedProject, nil
}
