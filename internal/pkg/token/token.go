package token

import (
	"errors"
	"os"
	"time"

	"github.com/Osas997/go-portfolio/internal/domains/auth/entity"
	"github.com/Osas997/go-portfolio/internal/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
)

type tokenResponse struct {
	AccessToken  string
	RefreshToken string
}

func signToken(payload *Payload, TOKEN_Key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString([]byte(TOKEN_Key))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func GenerateToken(user *entity.User) (*tokenResponse, error) {
	accessExpiryStr := os.Getenv("JWT_EXPIRED_AT")
	if accessExpiryStr == "" {
		accessExpiryStr = "10m"
	}

	accessExpiry, err := utils.ParseDuration(accessExpiryStr)
	if err != nil {
		return &tokenResponse{}, err
	}

	refreshExpiryStr := os.Getenv("JWT_REFRESH_EXPIRED_AT")
	if refreshExpiryStr == "" {
		refreshExpiryStr = "7d"
	}

	refreshExpiry, err := utils.ParseDuration(refreshExpiryStr)
	if err != nil {
		return &tokenResponse{}, err
	}

	accessTokenPayload := &Payload{
		Sub:      user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accesToken, err := signToken(accessTokenPayload, os.Getenv("JWT_SECRET"))
	if err != nil {
		return &tokenResponse{}, err
	}

	refreshTokenPayload := &Payload{
		Sub:      user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	refreshToken, err := signToken(refreshTokenPayload, os.Getenv("JWT_REFRESH_SECRET"))
	if err != nil {
		return &tokenResponse{}, err
	}

	return &tokenResponse{AccessToken: accesToken, RefreshToken: refreshToken}, nil
}

func VerifyToken(tokenStr string, secret string) (*Payload, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		// Pastikan metode signing benar
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Payload)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
