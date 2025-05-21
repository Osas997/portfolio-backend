package repository

import "github.com/Osas997/go-portfolio/internal/domains/projects/entity"

type ProjectRepository interface {
	Save(project *entity.Projects) (*entity.Projects, error)
	FindAll() ([]entity.Projects, error)
	FindById(id string) (*entity.Projects, error)
	Delete(id string) error
}
