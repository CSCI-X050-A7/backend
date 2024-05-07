package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Order struct to describe order object.
type Order struct {
	ID              uuid.UUID `gorm:"primarykey;type:uuid;default:(uuid_generate_v4())"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Tickets         []Ticket
	UserID          uuid.UUID
	PromotionID     uuid.UUID
	ShowID          uuid.UUID
	CardID          uuid.UUID
	MovieTitle      string
	TicketPrice     float64
	BookingFeePrice float64
	PromotionPrice  float64
	SalesTaxPrice   float64
	TotalPrice      float64
	CheckOut        bool
}
