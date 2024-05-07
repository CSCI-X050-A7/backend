package schema

import (
	"github.com/google/uuid"
)

type Ticket struct {
	ID      uuid.UUID `json:"id"`
	OrderID uuid.UUID `json:"order_id" validate:"required"`
	Seat    string    `json:"seat" validate:"required"`
	ShowID  uuid.UUID `json:"show_id" validate:"required"`
	Type    string    `json:"type" validate:"required"`
	Price   float64   `json:"price" validate:"required"`
}

type UpsertTicket struct {
	OrderID uuid.UUID `json:"order_id" validate:"required"`
	Seat    string    `json:"seat" validate:"required"`
	ShowID  uuid.UUID `json:"show_id" validate:"required"`
	Type    string    `json:"type" validate:"required"`
	Price   float64   `json:"price" validate:"required"`
}
