package booking

import (
	"github.com/google/uuid"
	"time"
)

func isValidBookingStatus(status BookingStatus) bool {
	switch status {
	case StatusPendingReview, StatusReserved, StatusRejected, StatusCancelled:
		return true
	default:
		return false
	}
}

type BookingStatus string

const (
	StatusPendingReview BookingStatus = "pending_review"
	StatusReserved      BookingStatus = "reserved"
	StatusRejected      BookingStatus = "rejected"
	StatusCancelled     BookingStatus = "cancelled"
)

type Booking struct {
	ID             uuid.UUID     `gorm:"type:uuid;primaryKey"`
	UserID         uuid.UUID     `gorm:"type:uuid;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OrganizationID uuid.UUID     `gorm:"type:uuid;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SpaceID        uuid.UUID     `gorm:"type:uuid;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StartTime      time.Time     `gorm:"not null"`
	EndTime        time.Time     `gorm:"not null"`
	Status         BookingStatus `gorm:"type:varchar(9);not null;default:'pending_review'"`
	Description    string        `gorm:"type:varchar(255);not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type BookingFilter struct {
	OrganizationID *uuid.UUID
	SpaceID        *uuid.UUID
	UserID         *uuid.UUID
	Status         *BookingStatus
	From           *time.Time
	To             *time.Time
}
