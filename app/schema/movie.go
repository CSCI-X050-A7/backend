package schema

import (
	"github.com/google/uuid"
)

// Movie struct to describe movie object.
type Movie struct {
	ID             uuid.UUID `json:"id" validate:"required"`
	Title          string    `json:"title" validate:"required,max=255"`
	Category       string    `json:"category" validate:"required,max=255"`
	Cast           string    `json:"cast" validate:"required,max=255"`
	Director       string    `json:"director" validate:"required,max=255"`
	Producer       string    `json:"producer" validate:"required,max=255"`
	Synopsis       string    `json:"synopsis" validate:"required,max=255"`
	Reviews        string    `json:"reviews" validate:"required,max=255"`
	TrailerPicture string    `json:"trailer_picture" validate:"required,max=255"`
	TrailerVideo   string    `json:"trailer_video" validate:"required,max=255"`
	RatingCode     string    `json:"rating_code" validate:"required,max=255"`
	ShowTime       string    `json:"show_time" validate:"required"`
}

type UpsertMovie struct {
	Title          string `json:"title" validate:"required,max=255"`
	Category       string `json:"category" validate:"required,max=255"`
	Cast           string `json:"cast" validate:"required,max=255"`
	Director       string `json:"director" validate:"required,max=255"`
	Producer       string `json:"producer" validate:"required,max=255"`
	Synopsis       string `json:"synopsis" validate:"required,max=255"`
	Reviews        string `json:"reviews" validate:"required,max=255"`
	TrailerPicture string `json:"trailer_picture" validate:"required,max=255"`
	TrailerVideo   string `json:"trailer_video" validate:"required,max=255"`
	RatingCode     string `json:"rating_code" validate:"required,max=255"`
	ShowTime       string `json:"show_time" validate:"required"`
}

type MovieListResponse = ListResponse[Movie]
