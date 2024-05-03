package schema

import (
	"github.com/google/uuid"
)

type Ticket struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	Seat  string    `json:"seat"`
	Show  string    `json:"show"`
	Type  string    `json:"type"`
	Price float64   `json:"price"`
}

type UpsertTicket struct {
	Title string  `json:"title"`
	Seat  string  `json:"seat"`
	Show  string  `json:"show"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}
