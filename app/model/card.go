package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Card struct to describe card object.
type Card struct {
	ID         uuid.UUID `gorm:"primarykey;type:uuid;default:(uuid_generate_v4())"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	CardType   string
	CardNumber string
	CardExpiry string
	Address    string
	Address2   string
	City       string
	State      string
	Zip        string
}
