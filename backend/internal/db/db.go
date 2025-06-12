package db

import (
	"fmt"

	"github.com/jdluques/uni-space-booking/internal/associations"
	"github.com/jdluques/uni-space-booking/internal/booking"
	"github.com/jdluques/uni-space-booking/internal/organization"
	"github.com/jdluques/uni-space-booking/internal/space"
	"github.com/jdluques/uni-space-booking/internal/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}

	if err := Migrate(db,
		&booking.Booking{},
		&organization.Organization{},
		&space.Space{},
		&user.User{},
		&associations.UserBooking{},
		&associations.UserOrganization{},
	); err != nil {
		return nil, fmt.Errorf("failed to migrate models: %w", err)
	}

	return db, nil
}
