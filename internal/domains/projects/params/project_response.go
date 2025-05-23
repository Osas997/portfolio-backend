package params

import (
	"github.com/Osas997/go-portfolio/internal/domains/projects/entity"
	"github.com/google/uuid"
)

type ProjectDetailResponse struct {
	ID            uuid.UUID                `json:"id"`
	Title         string                   `json:"title"`
	Content       string                   `json:"content"`
	Thumbnail     string                   `json:"thumbnail"`
	ProjectImages []*ProjectImagesRepsonse `json:"project_images"`
}

type ProjectImagesRepsonse struct {
	ID    uuid.UUID `json:"id"`
	Image string    `json:"image"`
}

type ProjectResponse struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Thumbnail string    `json:"thumbnail"`
}

func NewProjectResponse(projects []*entity.Projects) []*ProjectResponse {
	projectResponse := []*ProjectResponse{}
	for _, project := range projects {
		projectResponse = append(projectResponse, &ProjectResponse{
			ID:        project.ID,
			Title:     project.Title,
			Content:   project.Content,
			Thumbnail: project.Thumbnail,
		})
	}
	return projectResponse
}

func NewProjectDetailResponse(project *entity.Projects) *ProjectDetailResponse {
	imagesResponse := []*ProjectImagesRepsonse{}
	for _, image := range project.ProjectImages {
		imagesResponse = append(imagesResponse, &ProjectImagesRepsonse{
			ID:    image.ID,
			Image: image.Image,
		})
	}

	projectResponse := &ProjectDetailResponse{
		ID:            project.ID,
		Title:         project.Title,
		Content:       project.Content,
		Thumbnail:     project.Thumbnail,
		ProjectImages: imagesResponse,
	}
	return projectResponse
}
