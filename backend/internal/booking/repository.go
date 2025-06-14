package booking

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type BookingRepository interface {
	FindBookingsByFilter(ctx context.Context, filter BookingFilter) ([]Booking, error)
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) FindBookingsByFilter(ctx context.Context, filter BookingFilter) ([]Booking, error) {
	db := r.db.WithContext(ctx).Model(&Booking{})

	if (filter.From != nil) != (filter.To != nil) {
		return nil, errors.New("both 'From' and 'To' must be provided together")
	}

	if filter.OrganizationID != nil {
		db = db.Where("organization_id = ?", *filter.OrganizationID)
	}

	if filter.SpaceID != nil {
		db = db.Where("space_id = ?", *filter.SpaceID)
	}

	if filter.UserID != nil {
		db = db.Where("user_id = ?", *filter.UserID)
	}

	if filter.Status != nil {
		db = db.Where("status = ?", *filter.Status)
	}

	if filter.From != nil && filter.To != nil {
		db = db.Where("start_time BETWEEN ? AND ?", *filter.From, *filter.To)
	}

	var bookings []Booking
	if err := db.Find(&bookings).Error; err != nil {
		return nil, fmt.Errorf("failed to find bookings with filter: %w", err)
	}

	return bookings, nil
}
