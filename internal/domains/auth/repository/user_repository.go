package repository

import "github.com/Osas997/go-portfolio/internal/domains/auth/entity"

type UserRepository interface {
	GetUserByUsername(username string) (*entity.User, error)
}
