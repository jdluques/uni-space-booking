package space

import (
	"github.com/google/uuid"
	"time"
)

type Space struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string    `gorm:"not null"`
	Capacity  int       `gorm:"not null"`
	Available bool      `gorm:"default:true;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
