package schema

import (
	"github.com/google/uuid"
)

// Show struct to describe movie object.
type Order struct {
	ID           uuid.UUID `json:"id" validate:"required"`
	TicketsArray string    `json:"tickets_array" validate:"required,lte=255"`
	Promotion    Promotion `json:"promotion"`
	Show         Show      `json:"show" validate:"required"`
	Card         Card      `json:"card"`
	Seats        []string  `json:"seats" validate:"required,lte=255"`
}

type UpsertOrder struct {
	TicketsArray string    `json:"tickets_array" validate:"required,lte=255"`
	Promotion    Promotion `json:"promotion"`
	Show         Show      `json:"show" validate:"required"`
	Card         Card      `json:"card"`
	Seats        []string  `json:"seats" validate:"required,lte=255"`
}

type OrderListResponse = ListResponse[Order]
