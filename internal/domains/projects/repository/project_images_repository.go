package repository

import "github.com/Osas997/go-portfolio/internal/domains/projects/entity"

type ProjectImagesRepository interface {
	Save(projectImages *entity.ProjectImages) (*entity.ProjectImages, error)
	FindAll() ([]entity.ProjectImages, error)
	FindById(id string) (*entity.ProjectImages, error)
	Delete(id string) error
	DeleteAllByProjectId(id string) error
}
