package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username      string    `json:"username" gorm:"unique,type:varchar(255)"`
	Password      string    `json:"password" gorm:"type:varchar(255)"`
	Refresh_token string    `json:"refresh_token" gorm:"type:text"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
