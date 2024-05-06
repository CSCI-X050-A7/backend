package schema

import (
	"github.com/google/uuid"
)

// Promotion struct to describe promotion object.
type Promotion struct {
	ID            uuid.UUID `json:"id" validate:"required"`
	IsExpired     bool      `json:"is_expired"`
	ExpiryDate    string    `json:"expiry_date" validate:"required"`
	MovieAffected string    `json:"movie_affected" validate:"required,max=255"`
	Discount      float64   `json:"discount"`
	Title         string    `json:"title" validate:"required,max=255"`
	Description   string    `json:"description" validate:"required,max=255"`
}

type UpsertPromotion struct {
	IsExpired     bool    `json:"is_expired" validate:"required"`
	ExpiryDate    string  `json:"expiry_date"`
	MovieAffected string  `json:"movie_affected" validate:"required,max=255"`
	Discount      float64 `json:"discount"`
	Title         string  `json:"title" validate:"required,max=255"`
	Description   string  `json:"description" validate:"required,max=255"`
}
type PromoListResponse = ListResponse[Promotion]
