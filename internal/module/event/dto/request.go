package dto

import "time"

type CreateRequest struct {
	Title        string    `json:"title" validate:"required,min=2,max=150"`
	Description  string    `json:"description" validate:"omitempty,max=1000"`
	Location     string    `json:"location" validate:"required"`
	StartsAt     time.Time `json:"starts_at" validate:"required"`
	TotalTickets int       `json:"total_tickets" validate:"required,gt=0"`
	Price        int       `json:"price" validate:"required,gte=0"`
}

type UpdateRequest struct {
	Title        string    `json:"title" validate:"omitempty,min=2,max=150"`
	Description  string    `json:"description" validate:"omitempty,max=1000"`
	Location     string    `json:"location" validate:"omitempty"`
	StartsAt     time.Time `json:"starts_at" validate:"omitempty"`
	TotalTickets int       `json:"total_tickets" validate:"omitempty,gt=0"`
	Price        int       `json:"price" validate:"omitempty,gte=0"`
}
