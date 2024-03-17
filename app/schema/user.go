package schema

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
	IsAdmin   bool      `json:"is_admin"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
}

type Auth struct {
	Username string `json:"username" default:"demo"`
	Password string `json:"password" default:"123456"`
}

type RegisterUser struct {
	UserName       string `validate:"required,gte=3,lte=50"`
	Email          string `validate:"required,email,lte=150"`
	Password       string `validate:"required,lte=100,gte=8"`
	Name           string `validate:"required,lte=100"`
	Phone          string `validate:"required,lte=20"`
	Address        string `validate:"required,lte=150"`
	Address2       string `validate:"lte=150"`
	City           string `validate:"required,lte=100"`
	State          string `validate:"required,lte=100"`
	Zip            string `validate:"required,lte=20"`
	CardType       string `validate:"lte=50"`
	CardNumber     string `validate:"lte=50"`
	CardExpiration string `validate:"lte=50"`
	CardAddress    string `validate:"lte=150"`
	CardAddress2   string `validate:"lte=150"`
	CardCity       string `validate:"lte=100"`
	CardState      string `validate:"lte=100"`
	CardZip        string `validate:"lte=20"`
	NeedPromotion  bool
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

type JWT struct {
	UserID uuid.UUID `json:"user_id"`
	Admin  bool      `json:"admin"`
	Exp    int64     `json:"exp"`
}
