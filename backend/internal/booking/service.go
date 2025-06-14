package booking

import (
	"context"
	"github.com/google/uuid"
)

type BookingService interface {
	GetBookingsByFilter(ctx context.Context, spaceID uuid.UUID) ([]Booking, error)
	GetBookingsByUser(ctx context.Context, userID uuid.UUID) ([]Booking, error)
}

type bookingService struct {
	bookingRepository BookingRepository
}

func NewBookingService(bookingRepository BookingRepository) BookingService {
	return &bookingService{
		bookingRepository: bookingRepository,
	}
}

func (s *bookingService) GetBookingsByFilter(ctx context.Context, spaceID uuid.UUID) ([]Booking, error) {
	status := StatusReserved
	filter := BookingFilter{
		SpaceID: &spaceID,
		Status:  &status,
	}

	bookings, err := s.bookingRepository.FindBookingsByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}

	return bookings, nil
}

func (s *bookingService) GetBookingsByUser(ctx context.Context, userID uuid.UUID) ([]Booking, error) {
	filter := BookingFilter{
		UserID: &userID,
	}

	bookings, err := s.bookingRepository.FindBookingsByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}

	return bookings, nil
}
