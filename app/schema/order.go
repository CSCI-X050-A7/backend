package schema

import (
	"github.com/google/uuid"
)

// Show struct to describe movie object.
type Order struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	Tickets     []Ticket  `json:"tickets" validate:"required"`
	UserID      uuid.UUID `json:"user_id" validate:"required"`
	PromotionID uuid.UUID `json:"promotion_id" validate:"required"`
	ShowID      uuid.UUID `json:"show_id" validate:"required"`
	CardID      uuid.UUID `json:"card_id" validate:"required"`
}

type UpsertOrder struct {
	Tickets     []Ticket  `json:"tickets" validate:"required"`
	UserID      uuid.UUID `json:"user_id" validate:"required"`
	PromotionID uuid.UUID `json:"promotion_id" validate:"required"`
	ShowID      uuid.UUID `json:"show_id" validate:"required"`
	CardID      uuid.UUID `json:"card_id" validate:"required"`
}

type OrderListResponse = ListResponse[Order]
