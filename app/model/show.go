package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Show  struct to describe show object.
type Show struct {
	ID                uuid.UUID `gorm:"primarykey;type:uuid;default:(uuid_generate_v4())" validate:"required"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	MovieID           uuid.UUID
	StartTime         time.Time
	EndTime           time.Time
	BookingFee        float64
	AdultTicketPrice  float64
	ChildTicketPrice  float64
	SeniorTicketPrice float64
	TheaterLocation   string
}
