package schema

import (
	"time"

	"github.com/google/uuid"
)

type Show struct {
	ID                uuid.UUID `json:"id" validate:"required"`
	Movie             Movie     `json:"movie" validate:"required,lte=255"`
	StartTime         time.Time `json:"start_time" validate:"required,lte=255"`
	EndTime           time.Time `json:"end_time" validate:"required,lte=255"`
	BookingFee        float32   `json:"booking_fee" validate:"required,lte=255"`
	AdultTicketPrice  float32   `json:"adult_ticket_price" validate:"required,lte=255"`
	ChildTicketPrice  float32   `json:"child_ticket_price" validate:"required,lte=255"`
	SeniorTicketPrice float32   `json:"senior_ticket_price" validate:"required,lte=255"`
	TicketType        string    `json:"ticket_type" validate:"required,lte=1023"`
	TheaterLocation   string    `json:"theater_location" validate:"required,lte=1023"`
}
type UpsertShow struct {
	Movie             Movie     `json:"movie" validate:"required,lte=255"`
	StartTime         time.Time `json:"start_time" validate:"required,lte=255"`
	EndTime           time.Time `json:"end_time" validate:"required,lte=255"`
	BookingFee        float32   `json:"booking_fee" validate:"required,lte=255"`
	AdultTicketPrice  float32   `json:"adult_ticket_price" validate:"required,lte=255"`
	ChildTicketPrice  float32   `json:"child_ticket_price" validate:"required,lte=255"`
	SeniorTicketPrice float32   `json:"senior_ticket_price" validate:"required,lte=255"`
	TicketType        string    `json:"ticket_type" validate:"required,lte=1023"`
	TheaterLocation   string    `json:"theater_location" validate:"required,lte=1023"`
}
type ShowListResponse = ListResponse[Order]
