package schema

import (
	"time"

	"github.com/google/uuid"
)

// Show struct to describe movie object.
type Order struct {
	ID              uuid.UUID `json:"id" validate:"required"`
	CreatedAt       time.Time `json:"created_at" validate:"required"`
	Tickets         []Ticket  `json:"tickets" validate:"required"`
	UserID          uuid.UUID `json:"user_id" validate:"required"`
	PromotionID     uuid.UUID `json:"promotion_id" validate:"required"`
	ShowID          uuid.UUID `json:"show_id" validate:"required"`
	CardID          uuid.UUID `json:"card_id" validate:"required"`
	MovieTitle      string    `json:"movie_title" validate:"required,lte=255"`
	TicketPrice     float64   `json:"ticket_price" validate:"required"`
	BookingFeePrice float64   `json:"booking_fee_price" validate:"required"`
	PromotionPrice  float64   `json:"promotion_price" validate:"required"`
	SalesTaxPrice   float64   `json:"sales_tax_price" validate:"required"`
	TotalPrice      float64   `json:"total_price" validate:"required"`
	CheckOut        bool      `json:"check_out" validate:"required"`
}

type CreateOrder struct {
	ShowID        uuid.UUID      `json:"show_id" validate:"required"`
	PromotionCode string         `json:"promotion_code" validate:"lte=255"`
	Tickets       []CreateTicket `json:"tickets" validate:"required"`
}

type OrderListResponse = ListResponse[Order]
