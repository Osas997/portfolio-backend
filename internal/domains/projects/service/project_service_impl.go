package service

import "github.com/Osas997/go-portfolio/internal/domains/projects/repository"

type ProjectServiceImpl struct {
	repository repository.ProjectRepository
}

func NewProjectService(repository repository.ProjectRepository) ProjectService {
	return &ProjectServiceImpl{
		repository,
	}
}

// Create implements ProjectService.
func (p *ProjectServiceImpl) Create() {
	panic("unimplemented")
}

// Delete implements ProjectService.
func (p *ProjectServiceImpl) Delete(userId string) {
	panic("unimplemented")
}

// FindAll implements ProjectService.
func (p *ProjectServiceImpl) FindAll() {
	panic("unimplemented")
}

// FindById implements ProjectService.
func (p *ProjectServiceImpl) FindById(userId string) {
	panic("unimplemented")
}

// Update implements ProjectService.
func (p *ProjectServiceImpl) Update(userId string) {
	panic("unimplemented")
}
