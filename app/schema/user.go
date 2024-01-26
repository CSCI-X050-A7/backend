package schema

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
	IsAdmin   bool      `json:"is_admin"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}

type Auth struct {
	Username string `json:"username" default:"demo"`
	Password string `json:"password" default:"123456"`
}

type CreateUser struct {
	IsAdmin   bool   `json:"is_admin"`
	IsActive  bool   `json:"is_active"`
	UserName  string `json:"username" validate:"required,lte=50,gte=5"`
	Email     string `json:"email" validate:"required,email,lte=150"`
	Password  string `json:"password" validate:"required,lte=100,gte=10"`
	FirstName string `json:"first_name" validate:"required,lte=100"`
	LastName  string `json:"last_name" validate:"required,lte=100"`
}

type UpdateUser struct {
	IsAdmin   bool   `json:"is_admin"`
	IsActive  bool   `json:"is_active"`
	FirstName string `json:"first_name" validate:"required,lte=100"`
	LastName  string `json:"last_name" validate:"required,lte=100"`
}

type UserListResponse = ListResponse[User]
