package schema

import (
	"github.com/google/uuid"
)

// Card struct to describe card object.
type Card struct {
	ID         uuid.UUID `json:"id" validate:"required"`
	Type       string    `json:"type" validate:"required,lte=255"`
	Number     string    `json:"number" validate:"required,lte=255"`
	Expiration string    `json:"expiration" validate:"required,lte=255"`
	Address    string    `json:"address" validate:"required,lte=255"`
	Address2   string    `json:"address2" validate:"required,lte=255"`
	City       string    `json:"city" validate:"required,lte=255"`
	State      string    `json:"state" validate:"required,lte=1023"`
	Zip        string    `json:"zip" validate:"required,lte=1023"`
}

type UpdateCard struct {
	Type       string `json:"type" validate:"required,lte=255"`
	Number     string `json:"number" validate:"required,lte=255"`
	Expiration string `json:"expiration" validate:"required,lte=255"`
	Address    string `json:"address" validate:"required,lte=255"`
	Address2   string `json:"address2" validate:"required,lte=255"`
	City       string `json:"city" validate:"required,lte=255"`
	State      string `json:"state" validate:"required,lte=1023"`
	Zip        string `json:"zip" validate:"required,lte=1023"`
}
