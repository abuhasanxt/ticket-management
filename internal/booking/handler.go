package booking

import (
	"errors"
	"gotickets/internal/event"
	httpresponse "gotickets/internal/httpResponse"
	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(s *service) *handler {
	return &handler{service: s}

}

func gerCurrentUserById(c *echo.Context) (uint, bool) {

	userId, ok := c.Get("user_id").(uint)
	return userId, ok
}

func bookingErrorResponse(c *echo.Context, err error) error {
	if errors.Is(err, ErrBookingNotFound) {
		return c.JSON(http.StatusNotFound, httpresponse.Error{
			Code:    http.StatusNotFound,
			Message: "Booking not found",
		})
	}

	if errors.Is(err, event.ErrEventNotFound) {
		return c.JSON(http.StatusNotFound, httpresponse.Error{
			Code:    http.StatusNotFound,
			Message: "Event not found",
			Details: err.Error(),
		})
	}

	if errors.Is(err, ErrNotEnoughTickets) {
		return c.JSON(http.StatusConflict, httpresponse.Error{
			Code:    http.StatusConflict,
			Message: "Not enough tickets available",
			Details: err.Error(),
		})
	}

	if errors.Is(err, ErrBookingAlreadyCancelled) {
		return c.JSON(http.StatusConflict, httpresponse.Error{
			Code:    http.StatusConflict,
			Message: "Booking is already cancelled",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusInternalServerError, httpresponse.Error{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
		Details: err.Error(),
	})
}
