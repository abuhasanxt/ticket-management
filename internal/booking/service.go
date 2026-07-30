package booking

import (
	"errors"
	"gotickets/internal/booking/dto"
	"gotickets/internal/event"

	"github.com/google/uuid"
)

type service struct {
	bookingRepo Repository
	eventRepo   event.Repository
}

func NewService(bookingRepo Repository, eventRepo event.Repository) *service {
	return &service{
		bookingRepo: bookingRepo,
		eventRepo:   eventRepo,
	}
}

func generateBookingCode() string {
	return "GT-" + uuid.New().String()
}

func (s *service) CreateBooking(userId uint, req dto.CreateRequest) (*dto.Response, error) {

//transaction and locking system

booking,err:=s.bookingRepo.CreateWithTicketUpdate(userId,req.EventID,req.Quantity)

if err!=nil {
	return nil,err
}

return booking.ToResponse(),err



//normal booking system
	// event, err := s.eventRepo.GetByID(req.EventID)

	// if err != nil {
	// 	return nil, err
	// }
	// if event.AvailableTickets < req.Quantity {
	// 	return nil, ErrNotEnoughTickets
	// }

	// booking := &Booking{
	// 	UserID:      userId,
	// 	EventID:     req.EventID,
	// 	Quantity:    req.Quantity,
	// 	Status:      BookingConfirmed,
	// 	TotalPrice:  req.Quantity * event.Price,
	// 	BookingCode: generateBookingCode(),
	// }

	// if err := s.bookingRepo.Create(booking); err != nil {
	// 	return nil, err
	// }

	// event.AvailableTickets = event.AvailableTickets - req.Quantity

	// if err := s.eventRepo.Update(event); err != nil {
	// 	return nil, err
	// }

	// return booking.ToResponse(), nil
}

func (s *service) GetMyBookings(userId uint) ([]*dto.Response, error) {

	bookings, err := s.bookingRepo.GetByUserID(userId)
	if err != nil {
		return nil, err
	}

	response := make([]*dto.Response, len(bookings))

	for i, b := range bookings {
		response[i] = b.ToResponse()
	}
	return response, nil

}

func (s *service) GetBookingByID(bookingId uint) (*dto.Response, error) {
	booking, err := s.bookingRepo.GetByID(bookingId)

	if err != nil {
		return nil, err

	}
	return booking.ToResponse(), err
}


func (s *service) CancelledRequest(bookingID uint, userID uint) (*dto.Response, error) {

    booking, err := s.bookingRepo.GetByIDAndUserID(bookingID, userID)
    if err != nil {
        return nil, errors.New("booking not found")
    }

    if booking.Status == BookingCancelled {
        return nil, errors.New("booking already cancelled")
    }

    booking.Status = BookingCancelled

    if err := s.bookingRepo.Update(booking); err != nil {
        return nil, err
    }

    event, err := s.eventRepo.GetByID(booking.EventID)
    if err != nil {
        return nil, err
    }

    event.AvailableTickets += booking.Quantity

    if err := s.eventRepo.Update(event); err != nil {
        return nil, err
    }

    return booking.ToResponse(), nil
}