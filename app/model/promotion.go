package model

import (
	"time"

	"github.com/google/uuid"
)

// User struct to describe User object.
type Promotion struct {
	ID         uuid.UUID `gorm:"primarykey;type:uuid;default:(uuid_generate_v4())"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	IsExpired  bool
	ExpiryDate string
	Discount   float64
}
