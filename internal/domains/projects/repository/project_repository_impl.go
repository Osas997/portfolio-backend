package repository

import (
	"github.com/Osas997/go-portfolio/internal/domains/projects/entity"
	"gorm.io/gorm"
)

type ProjectRepositoryImpl struct {
	DB *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &ProjectRepositoryImpl{DB: db}
}

// Delete implements ProjectRepository.
func (p *ProjectRepositoryImpl) Delete(id string) error {
	panic("unimplemented")
}

// FindAll implements ProjectRepository.
func (p *ProjectRepositoryImpl) FindAll() ([]entity.Projects, error) {
	panic("unimplemented")
}

// FindById implements ProjectRepository.
func (p *ProjectRepositoryImpl) FindById(id string) (*entity.Projects, error) {
	panic("unimplemented")
}

// Save implements ProjectRepository.
func (p *ProjectRepositoryImpl) Save(project *entity.Projects) (*entity.Projects, error) {
	panic("unimplemented")
}
