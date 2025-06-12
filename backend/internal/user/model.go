package user

import (
	"github.com/google/uuid"
	"time"
)

type UserRole string

const (
	RoleStudent UserRole = "student"
	RoleTeacher UserRole = "teacher"
	RoleAdmin   UserRole = "admin"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	GoogleID  string    `gorm:"type:varchar(50);uniqueIndex;not null"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Role      UserRole  `gorm:"type:varchar(7);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
