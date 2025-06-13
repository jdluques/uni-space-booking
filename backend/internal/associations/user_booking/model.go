package user_booking

import "github.com/google/uuid"

type ApprovalStatus string

const (
	PendingAssignment ApprovalStatus = "pending_assignment"
	UnderReview       ApprovalStatus = "under_review"
	Reviewed          ApprovalStatus = "reviewed"
)

type UserBooking struct {
	RequesterID    uuid.UUID      `gorm:"type:uuid;primaryKey"`
	BookingID      uuid.UUID      `gorm:"type:uuid;primaryKey"`
	ApproverID     uuid.UUID      `gorm:"type:uuid"`
	ApprovalStatus ApprovalStatus `gorm:"type:varchar(18);not null;default:'pending_assignment'"`
}
