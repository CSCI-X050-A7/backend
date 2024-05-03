package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	// "gorm.io/gorm"
)

type Order struct {
	ID        uuid.UUID `gorm:"primarykey;type:uuid;default:(uuid_generate_v4())"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Tickets   []Ticket
	Promotion Promotion
	Show      Show
	Card      Card
}
