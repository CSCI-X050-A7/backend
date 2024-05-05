package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Order struct to describe order object.
type Order struct {
	ID           uuid.UUID `gorm:"primarykey;type:uuid;default:(uuid_generate_v4())"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	TicketsArray []Ticket
	Promotion    Promotion
	Show         Show
	Card         Card
	Seats        []string
}
