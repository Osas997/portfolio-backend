package service

import (
	"github.com/Osas997/go-portfolio/internal/domains/auth/params"
	"github.com/Osas997/go-portfolio/internal/domains/auth/repository"
)

type AuthServiceImpl struct {
	repository repository.UserRepository
}

func NewAuthService(repository repository.UserRepository) AuthService {
	return &AuthServiceImpl{
		repository,
	}
}

func (a *AuthServiceImpl) Login(authRequest *params.AuthRequest) {

}
