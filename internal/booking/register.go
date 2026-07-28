package booking

import (
	"gotickets/internal/event"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	bookingRepo := NewRepository(db)
	eventRepo :=event.NewRepository(db)

	svc := NewService(bookingRepo,eventRepo)
	handler := NewHandler(svc)

	api := e.Group("/api/v1/bookings")

	api.POST("", handler.CreateBooking)
	// api.GET("", handler.GetEvents)
	// api.GET("/:id", handler.GetEventById)
	// api.PATCH("/:id", handler.UpdatedEvent)

}
