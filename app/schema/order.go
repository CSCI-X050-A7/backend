package schema

import "github.com/google/uuid"

type Order struct {
	ID    uuid.UUID `json:"id" validate:"required"`
	Field string    `json:"field" validate:"required"`
}

type OrderListResponse = ListResponse[Order]
