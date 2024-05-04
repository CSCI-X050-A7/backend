package schema

import (
	"github.com/google/uuid"
)

// Promotion struct to describe promotion object.
type Promotion struct {
	ID         uuid.UUID `json:"id" validate:"required"`
	IsExpired  bool      `json:"is_expired" validate:"required"`
	ExpiryDate string    `json:"expiry_date" validate:"required,lte=255"`
	Discount   float64   `json:"discount" validate:"required"`
}

type UpsertPromotion struct {
	IsExpired  bool    `json:"is_expired" validate:"required"`
	ExpiryDate string  `json:"expiry_date" validate:"required,lte=255"`
	Discount   float64 `json:"discount" validate:"required"`
}

type PromoListResponse = ListResponse[Promotion]
