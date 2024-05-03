package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Show  struct to describe show object.
type Show struct {
	ID                uuid.UUID `gorm:"primarykey;type:uuid;default:(uuid_generate_v4())"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	Movie             Movie
	StartTime         time.Time
	EndTime           time.Time
	BookingFee        float32
	AdultTicketPrice  float32
	ChildTicketPrice  float32
	SeniorTicketPrice float32
	TheaterLocation   string
}
