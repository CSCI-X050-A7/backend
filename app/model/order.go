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
	TicketsArray []Ticket       // replace with Ticket[] once that struct is created
	Promotion    Promotion
	Show         Show
	Card         Card
	Seats        []string
}
