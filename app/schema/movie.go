package schema

import (
	"time"

	"github.com/google/uuid"
)

// Movie struct to describe movie object.
type Movie struct {
	ID             uuid.UUID `json:"id" validate:"required"`
	Title          string    `json:"title" validate:"required,lte=255"`
	Category       string    `json:"category" validate:"required,lte=255"`
	Cast           string    `json:"cast" validate:"required,lte=255"`
	Director       string    `json:"director" validate:"required,lte=255"`
	Producer       string    `json:"producer" validate:"required,lte=255"`
	Synopsis       string    `json:"synopsis" validate:"required,lte=255"`
	Reviews        string    `json:"reviews" validate:"required,lte=255"`
	TrailerPicture string    `json:"trailer_picture" validate:"required,lte=1023"`
	TrailerVideo   string    `json:"trailer_video" validate:"required,lte=1023"`
	RatingCode     string    `json:"rating_code" validate:"required,lte=255"`
	ShowTime       time.Time `json:"show_time" validate:"required"`
}

type UpsertMovie struct {
	Title          string    `json:"title" validate:"required,lte=255"`
	Category       string    `json:"category" validate:"required,lte=255"`
	Cast           string    `json:"cast" validate:"required,lte=255"`
	Director       string    `json:"director" validate:"required,lte=255"`
	Producer       string    `json:"producer" validate:"required,lte=255"`
	Synopsis       string    `json:"synopsis" validate:"required,lte=255"`
	Reviews        string    `json:"reviews" validate:"required,lte=255"`
	TrailerPicture string    `json:"trailer_picture" validate:"required,lte=1023"`
	TrailerVideo   string    `json:"trailer_video" validate:"required,lte=1023"`
	RatingCode     string    `json:"rating_code" validate:"required,lte=255"`
	ShowTime       time.Time `json:"show_time" validate:"required"`
}

type MovieListResponse = ListResponse[Movie]
