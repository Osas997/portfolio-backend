package service

import "github.com/Osas997/go-portfolio/internal/domains/auth/params"

type AuthService interface {
	Login(authRequest *params.AuthRequest) (*params.AuthResponse, error)
	Refresh(refreshRequest *params.RefreshTokenRequest) (*params.RefreshTokenResponse, error)
	Logout(userId string) error
}
