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
	ID             uuid.UUID `json:"id" validate:"required"`
	IsActive       bool      `json:"is_active" validate:"required"`
	IsAdmin        bool      `json:"is_admin" validate:"required"`
	UserName       string    `json:"username" validate:"required"`
	Email          string    `json:"email" validate:"required"`
	Name           string    `json:"name" validate:"required"`
	Phone          string    `json:"phone" validate:"required"`
	Address        string    `json:"address" validate:"required"`
	Address2       string    `json:"address2" validate:"required"`
	City           string    `json:"city" validate:"required"`
	State          string    `json:"state" validate:"required"`
	Zip            string    `json:"zip" validate:"required"`
	CardType       string    `json:"card_type" validate:"required"`
	CardNumber     string    `json:"card_number" validate:"required"`
	CardExpiration string    `json:"card_expiration" validate:"required"`
	CardAddress    string    `json:"card_address" validate:"required"`
	CardAddress2   string    `json:"card_address2" validate:"required"`
	CardCity       string    `json:"card_city" validate:"required"`
	CardState      string    `json:"card_state" validate:"required"`
	CardZip        string    `json:"card_zip" validate:"required"`
	NeedPromotion  bool      `json:"need_promotion" validate:"required"`
}

type RegisterUser struct {
	UserName       string `json:"username" validate:"required,gte=3,lte=50"`
	Email          string `json:"email" validate:"required,email,lte=150"`
	Password       string `json:"password" validate:"required,lte=100,gte=8"`
	Name           string `json:"name" validate:"required,lte=100"`
	Phone          string `json:"phone" validate:"required,number,eq=10"`
	Address        string `json:"address" validate:"required,lte=150"`
	Address2       string `json:"address2" validate:"lte=150"`
	City           string `json:"city" validate:"required,lte=100"`
	State          string `json:"state" validate:"required,lte=100"`
	Zip            string `json:"zip" validate:"required,number,eq=5"`
	CardType       string `json:"card_type" validate:"lte=50"`
	CardNumber     string `json:"card_number" validate:"lte=50"`
	CardExpiration string `json:"card_expiration" validate:"lte=50"`
	CardAddress    string `json:"card_address" validate:"lte=150"`
	CardAddress2   string `json:"card_address2" validate:"lte=150"`
	CardCity       string `json:"card_city" validate:"lte=100"`
	CardState      string `json:"card_state" validate:"lte=100"`
	CardZip        string `json:"card_zip" validate:"number,eq=5"`
	NeedPromotion  bool   `json:"need_promotion"`
}

type UpdateUser struct {
	UserName       string `json:"username" validate:"gte=3,lte=50"`
	Name           string `json:"name" validate:"lte=100"`
	Phone          string `json:"phone" validate:"lte=20"`
	Address        string `json:"address" validate:"lte=150"`
	Address2       string `json:"address2" validate:"lte=150"`
	City           string `json:"city" validate:"lte=100"`
	State          string `json:"state" validate:"lte=100"`
	Zip            string `json:"zip" validate:"lte=20"`
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

type AdminUpdateUser struct {
	UserName       string `json:"username" validate:"gte=3,lte=50"`
	Name           string `json:"name" validate:"lte=100"`
	Phone          string `json:"phone" validate:"lte=20"`
	Address        string `json:"address" validate:"lte=150"`
	Address2       string `json:"address2" validate:"lte=150"`
	City           string `json:"city" validate:"lte=100"`
	State          string `json:"state" validate:"lte=100"`
	Zip            string `json:"zip" validate:"lte=20"`
	CardType       string `json:"card_type" validate:"lte=50"`
	CardNumber     string `json:"card_number" validate:"lte=50"`
	CardExpiration string `json:"card_expiration" validate:"lte=50"`
	CardAddress    string `json:"card_address" validate:"lte=150"`
	CardAddress2   string `json:"card_address2" validate:"lte=150"`
	CardCity       string `json:"card_city" validate:"lte=100"`
	CardState      string `json:"card_state" validate:"lte=100"`
	CardZip        string `json:"card_zip" validate:"lte=20"`
	NeedPromotion  bool   `json:"need_promotion"`
	Email          string `json:"email" validate:"email,lte=150"`
	IsActive       bool   `json:"is_active"`
	IsAdmin        bool   `json:"is_admin"`
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

type UserDetailListResponse = ListResponse[UserDetail]

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
