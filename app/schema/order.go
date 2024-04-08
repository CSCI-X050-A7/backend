package schema

import "github.com/google/uuid"

type Order struct {
	ID uuid.UUID `json:"id"`
	// TODO: more fields
}

type OrderListResponse = ListResponse[Order]
