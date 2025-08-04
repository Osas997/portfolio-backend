package service

import "github.com/Osas997/go-portfolio/internal/domains/auth/params"

type AuthService interface {
	Login(authRequest *params.AuthRequest) (*params.AuthResponse, error)
	Refresh(refreshToken string) (*params.RefreshTokenResponse, error)
	Logout(userId string) error
	CsrfToken() string
}
