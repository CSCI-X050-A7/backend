package schema

import (
	"github.com/google/uuid"
)

// Show struct to describe movie object.
type Order struct {
	ID           uuid.UUID `json:"id" validate:"required"`
	TicketsArray []Ticket  `json:"tickets_array"`
	PromotionID  uuid.UUID `json:"promotion_id"`
	Promotion    Promotion `json:"promotion"`
	ShowID       uuid.UUID `json:"show_id" validate:"required"`
	Show         Show      `json:"show" validate:"required"`
	CardID       uuid.UUID `json:"card_id"`
	Card         Card      `json:"card"`
	Seats        []string  `json:"seats" validate:"required,lte=255"`
}

type UpsertOrder struct {
	TicketsArray []Ticket  `json:"tickets_array"`
	PromotionID  uuid.UUID `json:"promotion_id"`
	Promotion    Promotion `json:"promotion"`
	ShowID       uuid.UUID `json:"show_id" validate:"required"`
	Show         Show      `json:"show" validate:"required"`
	CardID       uuid.UUID `json:"card_id"`
	Card         Card      `json:"card"`
	Seats        []string  `json:"seats" validate:"required,lte=255"`
}

type OrderListResponse = ListResponse[Order]
