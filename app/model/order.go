package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct to describe User object.
type Order struct {
	ID        uuid.UUID `gorm:"primarykey;type:uuid;default:(uuid_generate_v4())"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	IsExpired bool
}
