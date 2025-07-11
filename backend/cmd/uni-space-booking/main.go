package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	// "github.com/jdluques/uni-space-booking/internal/associations/user_booking"
	// "github.com/jdluques/uni-space-booking/internal/associations/user_organization"
	"github.com/jdluques/uni-space-booking/internal/booking"
	"github.com/jdluques/uni-space-booking/internal/db"
	// "github.com/jdluques/uni-space-booking/internal/organization"
	// "github.com/jdluques/uni-space-booking/internal/space"
	// "github.com/jdluques/uni-space-booking/internal/user"
)

func registerAllRoutes(e *echo.Echo, bookingHandler *booking.BookingHandler) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "OK",
		})
	})

	booking.RegisterBookingRoutes(e, bookingHandler)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		dbUser, dbPassword, dbName, dbHost, dbPort, dbSSLMode)

	db, err := db.Connect(dsn)
	if err != nil {
		log.Println("Error connecting to database:", err)
		return
	}
	fmt.Println("Successfully connected to the database!")

	bookingRepo := booking.NewBookingRepository(db)
	// organizationRepo := organization.NewOrganizationRepository(db)
	// spaceRepo := space.NewSpaceRepository(db)
	// userRepo := user.NewUserRepository(db)
	// userBookingRepo := user_booking.NewUserBookingRepository(db)
	// userOrganizationRepo := user_organization.NewUserOrganizationRepository(db)

	bookingService := booking.NewBookingService(bookingRepo)

	bookingHandler := booking.NewBookingHandler(bookingService)

	e := echo.New()

	registerAllRoutes(e, bookingHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
