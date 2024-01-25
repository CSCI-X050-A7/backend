package schema

import (
	"time"

	"github.com/google/uuid"
)

// Book struct to describe book object.
type Book struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id" validate:"required"`
	Title     string    `json:"title" validate:"required,lte=255"`
	Author    string    `json:"author" validate:"required,lte=255"`
	Status    int       `json:"status" validate:"required,len=1"`
	Meta      Meta      `gorm:"embedded" json:"meta" validate:"required"`
}

// Meta struct to describe book attributes.
type Meta struct {
	Picture     string `json:"picture"`
	Description string `json:"description"`
	Rating      int    `json:"rating" validate:"min=1,max=10"`
}

type CreateBook struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
	Title  string    `json:"title" validate:"required,lte=255"`
	Author string    `json:"author" validate:"required,lte=255"`
	Status int       `json:"status" validate:"required,len=1"`
	Meta   Meta      `gorm:"embedded" json:"meta" validate:"required"`
}

type UpdateBook struct {
	Title  string `json:"title" validate:"required,lte=255"`
	Author string `json:"author" validate:"required,lte=255"`
	Status int    `json:"status" validate:"required,len=1"`
	Meta   Meta   `gorm:"embedded" json:"meta" validate:"required"`
}
