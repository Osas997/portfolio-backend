package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username  string    `json:"username" gorm:"unique,type:varchar(255)"`
	Password  string    `json:"password" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// func (u *User) BeforeCreate(tx *gorm.DB) error {
// 	if u.ID == uuid.Nil {
// 		u.ID = uuid.New()
// 	}
// 	return nil
// }
