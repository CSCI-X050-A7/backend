package schema

import (
	"github.com/google/uuid"
)

type Show struct {
	ID                uuid.UUID `json:"id" validate:"required"`
	Movie             Movie     `json:"movie" validate:"required"`
	StartTime         string    `json:"start_time" validate:"required"`
	EndTime           string    `json:"end_time" validate:"required"`
	BookingFee        float32   `json:"booking_fee" validate:"required"`
	AdultTicketPrice  float32   `json:"adult_ticket_price" validate:"required"`
	ChildTicketPrice  float32   `json:"child_ticket_price" validate:"required"`
	SeniorTicketPrice float32   `json:"senior_ticket_price" validate:"required"`
	TicketType        string    `json:"ticket_type" validate:"required,max=1023"`
	TheaterLocation   string    `json:"theater_location" validate:"required,max=1023"`
}

type UpsertShow struct {
	Movie             Movie   `json:"movie" validate:"required"`
	StartTime         string  `json:"start_time" validate:"required"`
	EndTime           string  `json:"end_time" validate:"required"`
	BookingFee        float32 `json:"booking_fee" validate:"required"`
	AdultTicketPrice  float32 `json:"adult_ticket_price" validate:"required"`
	ChildTicketPrice  float32 `json:"child_ticket_price" validate:"required"`
	SeniorTicketPrice float32 `json:"senior_ticket_price" validate:"required"`
	TicketType        string  `json:"ticket_type" validate:"required,max=1023"`
	TheaterLocation   string  `json:"theater_location" validate:"required,max=1023"`
}

type ShowListResponse = ListResponse[Show]
