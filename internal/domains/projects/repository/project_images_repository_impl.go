package repository

import (
	"github.com/Osas997/go-portfolio/internal/domains/projects/entity"
	"gorm.io/gorm"
)

type ProjectImagesRepositoryImpl struct {
	DB *gorm.DB
}

func NewProjectImagesRepository(db *gorm.DB) ProjectImagesRepository {
	return &ProjectImagesRepositoryImpl{DB: db}
}

// Delete implements ProjectImagesRepository.
func (p *ProjectImagesRepositoryImpl) Delete(id string) error {
	panic("unimplemented")
}

// FindAll implements ProjectImagesRepository.
func (p *ProjectImagesRepositoryImpl) FindAll() ([]entity.ProjectImages, error) {
	panic("unimplemented")
}

// FindById implements ProjectImagesRepository.
func (p *ProjectImagesRepositoryImpl) FindById(id string) (*entity.ProjectImages, error) {
	panic("unimplemented")
}

func (p *ProjectImagesRepositoryImpl) DeleteAllByProjectId(id string) error {
	if err := p.DB.Where("project_id = ?", id).Delete(&entity.ProjectImages{}).Error; err != nil {
		return err
	}
	return nil
}

// Save implements ProjectImagesRepository.

func (p *ProjectImagesRepositoryImpl) Save(projectImages *entity.ProjectImages) (*entity.ProjectImages, error) {
	if err := p.DB.Save(projectImages).Error; err != nil {
		return &entity.ProjectImages{}, err
	}
	return projectImages, nil
}
