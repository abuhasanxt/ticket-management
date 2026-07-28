package booking

import (
	"errors"
	"gotickets/internal/booking/dto"
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

func getCurrentUserById(c *echo.Context) (uint, bool) {

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





func (h *handler)CreateBooking(c *echo.Context)error{
	userId ,ok :=getCurrentUserById(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized,httpresponse.Error{
			Code: http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}

	var req dto.CreateRequest
	if err:=c.Bind(&req);err!=nil {
		return c.JSON(http.StatusBadRequest,httpresponse.Error{
			Code: http.StatusBadRequest,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
		
	}


	if err:=c.Validate(&req);err!=nil {
		return  c.JSON(http.StatusBadRequest,httpresponse.Error{
			Code: http.StatusBadRequest,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}




	response ,err:=h.service.CreateBooking(userId,req)

	if err!=nil {
		return bookingErrorResponse(c,err)
	}

	return c.JSON(http.StatusCreated,response)
}