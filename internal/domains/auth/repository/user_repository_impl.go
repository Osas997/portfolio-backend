package repository

import (
	"github.com/Osas997/go-portfolio/internal/domains/auth/entity"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) GetUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetUserById(id string) (*entity.User, error) {
	var user entity.User
	err := r.DB.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Save(user *entity.User) (*entity.User, error) {
	if err := r.DB.Save(user).Error; err != nil {
		return &entity.User{}, err
	}
	return user, nil
}
