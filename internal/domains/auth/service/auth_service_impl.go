package service

import (
	"os"

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

	user.Refresh_token = tokens.RefreshToken
	if _, err := a.repository.Save(user); err != nil {
		return nil, err
	}

	return &params.AuthResponse{AccessToken: tokens.AccessToken, RefreshToken: tokens.RefreshToken}, nil
}

func (a *AuthServiceImpl) Logout(userId string) error {
	user, err := a.repository.GetUserById(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errorhandler.NewNotFoundError("User not found")
		}
		return err
	}

	user.Refresh_token = ""
	if _, err := a.repository.Save(user); err != nil {
		return err
	}
	return nil
}

func (a *AuthServiceImpl) Refresh(refreshRequest *params.RefreshTokenRequest) (*params.RefreshTokenResponse, error) {
	payload, err := token.VerifyToken(refreshRequest.RefreshToken, os.Getenv("JWT_REFRESH_SECRET"))
	if err != nil {
		return nil, errorhandler.NewNotFoundError("Invalid refresh token")
	}

	user, err := a.repository.GetUserById(payload.Sub.String())
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorhandler.NewNotFoundError("User not found")
		}
		return nil, err
	}

	if user.Refresh_token != refreshRequest.RefreshToken {
		return nil, errorhandler.NewNotFoundError("Invalid refresh token")
	}

	accessToken, err := token.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &params.RefreshTokenResponse{AccessToken: accessToken.AccessToken}, nil
}
