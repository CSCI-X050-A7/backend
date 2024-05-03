package schema

import "github.com/google/uuid"

type Order struct {
	ID uuid.UUID `json:"id" validate:"required"`
	// TODO: more fields
}

type OrderListResponse = ListResponse[Order]
