package booking

import (
	"context"
)

type BookingService interface {
	GetBookingsByFilter(ctx context.Context, filter BookingFilter) ([]Booking, error)
}

type bookingService struct {
	bookingRepository BookingRepository
}

func NewBookingService(bookingRepository BookingRepository) BookingService {
	return &bookingService{
		bookingRepository: bookingRepository,
	}
}

func (s *bookingService) GetBookingsByFilter(ctx context.Context, filter BookingFilter) ([]Booking, error) {
	bookings, err := s.bookingRepository.FindBookingsByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	return bookings, nil
}
