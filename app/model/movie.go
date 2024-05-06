package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Movie struct to describe movie object.
type Movie struct {
	ID             uuid.UUID `gorm:"primarykey;type:uuid;default:(uuid_generate_v4())"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Title          string
	Category       string
	Cast           string
	Director       string
	Producer       string
	Synopsis       string
	Reviews        string
	TrailerPicture string
	TrailerVideo   string
	RatingCode     string
	ShowTime       time.Time
}
