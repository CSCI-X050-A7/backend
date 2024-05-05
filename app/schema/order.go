package schema

import (
	"github.com/google/uuid"
)

// Show struct to describe movie object.
type Order struct {
	ID           uuid.UUID `json:"id" validate:"required"`
	TicketsArray string    `json:"tickets_array" validate:"required,lte=255"`
	Promotion    Promotion `json:"promotion" validate:"required,lte=255"`
	Show         Show      `json:"show" validate:"required,lte=255"`
	Card         string    `json:"card"`
	Seats        []string  `json:"seats" validate:"required,lte=255"`
}

// Show struct to describe movie object.
type UpsertOrder struct {
	TicketsArray string    `json:"tickets_array" validate:"required,lte=255"`
	Promotion    Promotion `json:"promotion" validate:"required,lte=255"`
	Show         Show      `json:"show" validate:"required,lte=255"`
	Card         string    `json:"card"`
	Seats        []string  `json:"seats" validate:"required,lte=255"`
}

type OrderListResponse = ListResponse[Order]
