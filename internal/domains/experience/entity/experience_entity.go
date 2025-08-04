package entity

import (
	"time"

	"github.com/google/uuid"
)

type Experience struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title     string    ` gorm:"type:varchar(255);not null"`
	Content   string    ` gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
