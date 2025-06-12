package associations

import "github.com/google/uuid"

type UserOrganization struct {
	UserID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;primaryKey"`
}
