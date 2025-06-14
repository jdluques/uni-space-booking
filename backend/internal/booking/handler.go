package booking

import (
	"fmt"
	"github.com/jdluques/uni-space-booking/internal/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type BookingHandler struct {
	bookingService BookingService
}

func NewBookingHandler(bookingService BookingService) *BookingHandler {
	return &BookingHandler{bookingService}
}

func (h *BookingHandler) GetReservedBookingsBySpace(c echo.Context) error {
	spaceID, err := utils.ParseUUIDParam(c, "space-id")
	if err != nil {
		return err
	}

	approvedBookingsBySpace, err := h.bookingService.GetBookingsByFilter(c.Request().Context(), spaceID)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			fmt.Sprintf("failed to fetch approved bookings of space %s ", spaceID),
		)
	}

	return c.JSON(http.StatusOK, approvedBookingsBySpace)
}

func (h *BookingHandler) GetBookingsByUser(c echo.Context) error {
	userID, err := utils.ParseUUIDParam(c, "user-id")
	if err != nil {
		return err
	}

	bookingsByUser, err := h.bookingService.GetBookingsByFilter(c.Request().Context(), userID)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			fmt.Sprintf("failed to fetch bookings of user %s", userID),
		)
	}

	return c.JSON(http.StatusOK, bookingsByUser)
}
