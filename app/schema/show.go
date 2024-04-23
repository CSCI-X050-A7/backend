package schema

import "github.com/google/uuid"

type Show struct {
	ID    uuid.UUID `json:"id" validate:"required"`
	Field string    `json:"field" validate:"required"`
}

type ShowListResponse = ListResponse[Order]
