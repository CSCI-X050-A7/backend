package schema

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" validate:"required"`
	IsActive bool      `json:"is_active"`
	IsAdmin  bool      `json:"is_admin"`
	UserName string    `json:"username"`
	Email    string    `json:"email"`
	Name     string    `json:"name"`
}

type UserDetail struct {
	ID             uuid.UUID `json:"id"`
	IsActive       bool      `json:"is_active"`
	IsAdmin        bool      `json:"is_admin"`
	UserName       string    `json:"username"`
	Email          string    `json:"email"`
	Name           string    `json:"name"`
	Phone          string    `json:"phone"`
	Address        string    `json:"address"`
	Address2       string    `json:"address2"`
	City           string    `json:"city"`
	State          string    `json:"state"`
	Zip            string    `json:"zip"`
	CardType       string    `json:"card_type"`
	CardNumber     string    `json:"card_number"`
	CardExpiration string    `json:"card_expiration"`
	CardAddress    string    `json:"card_address"`
	CardAddress2   string    `json:"card_address2"`
	CardCity       string    `json:"card_city"`
	CardState      string    `json:"card_state"`
	CardZip        string    `json:"card_zip"`
	NeedPromotion  bool      `json:"need_promotion"`
}

type RegisterUser struct {
	UserName       string `json:"username" validate:"required,gte=3,lte=50"`
	Email          string `json:"email" validate:"required,email,lte=150"`
	Password       string `json:"password" validate:"required,lte=100,gte=8"`
	Name           string `json:"name" validate:"required,lte=100"`
	Phone          string `json:"phone" validate:"required,lte=20"`
	Address        string `json:"address" validate:"required,lte=150"`
	Address2       string `json:"address2" validate:"lte=150"`
	City           string `json:"city" validate:"required,lte=100"`
	State          string `json:"state" validate:"required,lte=100"`
	Zip            string `json:"zip" validate:"required,lte=20"`
	CardType       string `json:"card_type" validate:"lte=50"`
	CardNumber     string `json:"card_number" validate:"lte=50"`
	CardExpiration string `json:"card_expiration" validate:"lte=50"`
	CardAddress    string `json:"card_address" validate:"lte=150"`
	CardAddress2   string `json:"card_address2" validate:"lte=150"`
	CardCity       string `json:"card_city" validate:"lte=100"`
	CardState      string `json:"card_state" validate:"lte=100"`
	CardZip        string `json:"card_zip" validate:"lte=20"`
	NeedPromotion  bool   `json:"need_promotion"`
}

type UpdateUser struct {
	UserName       string `json:"username" validate:"required,gte=3,lte=50"`
	Password       string `json:"password" validate:"required,lte=100,gte=8"`
	Name           string `json:"name" validate:"required,lte=100"`
	Phone          string `json:"phone" validate:"required,lte=20"`
	Address        string `json:"address" validate:"required,lte=150"`
	Address2       string `json:"address2" validate:"lte=150"`
	City           string `json:"city" validate:"required,lte=100"`
	State          string `json:"state" validate:"required,lte=100"`
	Zip            string `json:"zip" validate:"required,lte=20"`
	CardType       string `json:"card_type" validate:"lte=50"`
	CardNumber     string `json:"card_number" validate:"lte=50"`
	CardExpiration string `json:"card_expiration" validate:"lte=50"`
	CardAddress    string `json:"card_address" validate:"lte=150"`
	CardAddress2   string `json:"card_address2" validate:"lte=150"`
	CardCity       string `json:"card_city" validate:"lte=100"`
	CardState      string `json:"card_state" validate:"lte=100"`
	CardZip        string `json:"card_zip" validate:"lte=20"`
	NeedPromotion  bool   `json:"need_promotion"`
}

type CreateUser struct {
	IsAdmin  bool   `json:"is_admin"`
	IsActive bool   `json:"is_active"`
	UserName string `json:"username" validate:"required,gte=3,lte=50"`
	Email    string `json:"email" validate:"required,email,lte=150"`
	Password string `json:"password" validate:"required,lte=100,gte=8"`
	Name     string `json:"name" validate:"required,lte=100"`
	Phone    string `json:"phone" validate:"required,lte=20"`
}

type UserListResponse = ListResponse[User]

type Auth struct {
	Username string `json:"username" default:"demo"`
	Password string `json:"password" default:"123456"`
	Remember bool   `json:"remember" default:"false"`
}

type JWT struct {
	UserID uuid.UUID `json:"user_id"`
	Admin  bool      `json:"admin"`
	Exp    int64     `json:"exp"`
}

type UserResetPassword struct {
	Username    string `json:"username"`
	NewPassword string `json:"newPassword"`
}

type UserChangePassword struct {
	Username        string `json:"username"`
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}
