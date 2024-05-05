package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Promotion struct to describe promotion object.
type Promotion struct {
	ID        uuid.UUID `gorm:"primarykey;type:uuid;default:(uuid_generate_v4())"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Discount  float64
	Code      string
}
