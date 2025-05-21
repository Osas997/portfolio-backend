package service

import (
	"github.com/Osas997/go-portfolio/internal/domains/auth/params"
	"github.com/Osas997/go-portfolio/internal/domains/auth/repository"
	"github.com/Osas997/go-portfolio/internal/pkg/errorhandler"
	"github.com/Osas997/go-portfolio/internal/pkg/hash"
	"github.com/Osas997/go-portfolio/internal/pkg/token"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	repository repository.UserRepository
}

func NewAuthService(repository repository.UserRepository) AuthService {
	return &AuthServiceImpl{
		repository,
	}
}

func (a *AuthServiceImpl) Login(authRequest *params.AuthRequest) (*params.AuthResponse, error) {
	user, err := a.repository.GetUserByUsername(authRequest.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorhandler.NewNotFoundError("Invalid username or password")
		}
		return nil, err
	}

	isPasswordValid := hash.ValidatePassword(user.Password, authRequest.Password)

	if !isPasswordValid {
		return nil, errorhandler.NewNotFoundError("Invalid username or password")
	}

	tokens, err := token.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &params.AuthResponse{AccessToken: tokens.AccessToken, RefreshToken: tokens.RefreshToken}, nil
}
