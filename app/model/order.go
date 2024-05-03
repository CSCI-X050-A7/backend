package model

import (
	"time"

	"github.com/google/uuid"
	// "gorm.io/gorm"
)

// Movie struct to describe movie object.
type Order struct {
	ID               uuid.UUID `gorm:"primarykey;type:uuid;default:(uuid_generate_v4())"`
	CreditCardNumber string
	MovieTitle       string
	ShowTime         time.Time
	TicketNumber     string
	BookingNumber    string
}
