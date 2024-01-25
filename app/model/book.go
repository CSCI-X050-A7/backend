package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Book struct to describe book object.
type Book struct {
	ID        uuid.UUID `gorm:"primarykey;type:uuid;default:(uuid_generate_v4())"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    uuid.UUID
	User      User
	Title     string
	Author    string
	Status    int
	Meta      Meta `gorm:"embedded"`
}

// Meta struct to describe book attributes.
type Meta struct {
	Picture     string
	Description string
	Rating      int
}
