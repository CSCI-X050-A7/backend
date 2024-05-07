package schema

import (
	"github.com/google/uuid"
)

// Promotion struct to describe promotion object.
type Promotion struct {
	ID       uuid.UUID `json:"id" validate:"required"`
	Discount float64   `json:"discount" validate:"required"`
	Code     string    `json:"code" validate:"required,lte=255"`
}

type UpsertPromotion struct {
	Discount float64 `json:"discount" validate:"required"`
	Code     string  `json:"code" validate:"required,lte=255"`
}
type PromoListResponse = ListResponse[Promotion]
