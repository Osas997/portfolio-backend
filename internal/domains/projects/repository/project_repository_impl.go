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
	if err := p.DB.Delete(&entity.Projects{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements ProjectRepository.
func (p *ProjectRepositoryImpl) FindAll() ([]*entity.Projects, error) {
	var projects []*entity.Projects
	if err := p.DB.Find(&projects).Order("created_at DESC").Error; err != nil {
		return nil, err
	}
	return projects, nil
}

// FindById implements ProjectRepository.
func (p *ProjectRepositoryImpl) FindById(id string) (*entity.Projects, error) {
	var project entity.Projects
	if err := p.DB.Preload("ProjectImages").First(&project, "id = ?", id).Error; err != nil {
		return &entity.Projects{}, err
	}
	return &project, nil
}

// Save implements ProjectRepository.
func (p *ProjectRepositoryImpl) Save(project *entity.Projects) (*entity.Projects, error) {
	if err := p.DB.Save(project).Error; err != nil {
		return &entity.Projects{}, err
	}
	return project, nil
}
