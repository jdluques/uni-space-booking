package organization

import (
	"github.com/google/uuid"
	"github.com/jdluques/uni-space-booking/internal/user"
	"time"
)

type Organization struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string    `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
