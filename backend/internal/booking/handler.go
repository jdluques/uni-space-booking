package booking

import (
	"fmt"
	"github.com/google/uuid"
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
	var filter BookingFilter

	spaceIDStr := c.Param("space")
	if spaceIDStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "space_id is required")
	}
	spaceId, err := uuid.Parse(spaceIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid space_id format")
	}

	status := StatusReserved

	filter = BookingFilter{Status: &status, SpaceID: &spaceId}

	approvedBookingsBySpace, err := h.bookingService.GetBookingsByFilter(c.Request().Context(), filter)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			fmt.Sprintf("failed to fetch approved bookings of %s space", spaceIDStr),
		)
	}

	return c.JSON(http.StatusOK, approvedBookingsBySpace)
}
