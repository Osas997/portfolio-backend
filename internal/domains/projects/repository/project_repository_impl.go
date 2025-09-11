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
	if err := p.DB.Order("created_at DESC").Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

// FindAllWithPagination returns a paginated list of projects and the total count
func (p *ProjectRepositoryImpl) FindAllWithPagination(page, limit int) ([]*entity.Projects, int64, error) {
	var (
		projects []*entity.Projects
		total    int64
	)
	// Count total
	if err := p.DB.Model(&entity.Projects{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit
	if err := p.DB.Order("created_at DESC").Limit(limit).Offset(offset).Find(&projects).Error; err != nil {
		return nil, 0, err
	}
	return projects, total, nil
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
