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
	Movie             Movie          `validate:"required"`
	StartTime         string         `validate:"required"`
	EndTime           string         `validate:"required"`
	BookingFee        float32        `validate:"required"`
	AdultTicketPrice  float32        `validate:"required"`
	ChildTicketPrice  float32        `validate:"required"`
	SeniorTicketPrice float32        `validate:"required"`
	TheaterLocation   string         `validate:"required"`
}
