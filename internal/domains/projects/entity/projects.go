package entity

import (
	"time"

	"github.com/google/uuid"
)

type Projects struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title     string    ` gorm:"type:varchar(255)"`
	Content   string    ` gorm:"type:text"`
	Thumbnail string    ` gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time

	ProjectImages []ProjectImages `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
