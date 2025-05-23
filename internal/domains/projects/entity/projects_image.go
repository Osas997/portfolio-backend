package entity

import (
	"time"

	"github.com/google/uuid"
)

type ProjectImages struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	ProjectID uuid.UUID `gorm:"type:uuid"`
	Image     string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
