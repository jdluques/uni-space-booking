package user_booking

import "gorm.io/gorm"

type UserBookingRepository interface {
}

type userBookingRepository struct {
	db *gorm.DB
}

func NewUserBookingRepository(db *gorm.DB) UserBookingRepository {
	return &userBookingRepository{db: db}
}
