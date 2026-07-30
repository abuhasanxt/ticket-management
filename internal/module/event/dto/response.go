package dto

import "time"

type Response struct {
	ID               uint      `json:"id"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	Location         string    `json:"location"`
	TotalTickets     int       `json:"total_tickets"`
	AvailableTickets int       `json:"available_tickets"`
	StartsAt         time.Time `json:"starts_at"`
	Price            int       `json:"price"`
	CreatedAt        string    `json:"created_at"`
}
