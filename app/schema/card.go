package schema

import (
	"github.com/google/uuid"
)

// Card struct to describe card object.
type Card struct {
	ID         uuid.UUID `json:"id" validate:"required"`
	CardType   string    `json:"card_type" validate:"required,lte=255"`
	CardNumber string    `json:"card_number" validate:"required,lte=255"`
	CardExpiry string    `json:"card_expiry" validate:"required,lte=255"`
	Address    string    `json:"address" validate:"required,lte=255"`
	Address2   string    `json:"address2" validate:"required,lte=255"`
	City       string    `json:"city" validate:"required,lte=255"`
	State      string    `json:"state" validate:"required,lte=1023"`
	Zip        string    `json:"zip" validate:"required,lte=1023"`
}
