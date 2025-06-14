package booking

import "github.com/labstack/echo/v4"

func RegisterBookingRoutes(e *echo.Echo, h *BookingHandler) {
	g := e.Group("/bookings")

	g.GET("/approved/:space_id", h.GetReservedBookingsBySpace)
	g.GET("/bookings/:user_id", h.GetBookingsByUser)
}
