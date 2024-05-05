package schema

import (
	"time"

	"github.com/google/uuid"
)

// Promotion struct to describe promotion object.
type Promotion struct {
	ID            uuid.UUID `json:"id" validate:"required"`
	IsExpired     bool      `json:"is_expired" validate:"required"`
	ExpiryDate    time.Time `json:"expiry_date" validate:"required,lte=255"`
	MovieAffected string    `json:"movie_affected" validate:"required,lte=255"`
	Discount      float64   `json:"discount" validate:"required"`
	Title         string    `json:"title" validate:"required,lte=255"`
	Description   string    `json:"description" validate:"required,lte=255"`
}

type UpsertPromotion struct {
	IsExpired     bool      `json:"is_expired" validate:"required"`
	ExpiryDate    time.Time `json:"expiry_date" validate:"required,lte=255"`
	MovieAffected string    `json:"movie_affected" validate:"required,lte=255"`
	Discount      float64   `json:"discount" validate:"required"`
	Title         string    `json:"title" validate:"required,lte=255"`
	Description   string    `json:"description" validate:"required,lte=255"`
}

type PromoListResponse = ListResponse[Promotion]
