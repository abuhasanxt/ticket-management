package dto

type CreateRequest struct {
	EventID  uint `json:"event_id"`
	Quantity int  `json:"quantity"`
}
