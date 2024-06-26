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
	ID            uuid.UUID `json:"id"`
	IsActive      bool      `json:"is_active"`
	IsAdmin       bool      `json:"is_admin"`
	UserName      string    `json:"username"`
	Email         string    `json:"email"`
	Name          string    `json:"name"`
	Phone         string    `json:"phone"`
	Address       string    `json:"address"`
	Address2      string    `json:"address2"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	Zip           string    `json:"zip"`
	NeedPromotion bool      `json:"need_promotion"`
	Cards         []Card    `json:"cards" validate:"required,max=3"`
}
type UserDetailNoCards struct {
	ID            uuid.UUID `json:"id"`
	IsActive      bool      `json:"is_active"`
	IsAdmin       bool      `json:"is_admin"`
	UserName      string    `json:"username"`
	Email         string    `json:"email"`
	Name          string    `json:"name"`
	Phone         string    `json:"phone"`
	Address       string    `json:"address"`
	Address2      string    `json:"address2"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	Zip           string    `json:"zip"`
	NeedPromotion bool      `json:"need_promotion"`
}

type RegisterUser struct {
	UserName      string       `json:"username" validate:"required,gte=3,lte=50"`
	Email         string       `json:"email" validate:"required,email,lte=150"`
	Password      string       `json:"password" validate:"required,lte=100,gte=8"`
	Name          string       `json:"name" validate:"required,lte=100"`
	Phone         string       `json:"phone" validate:"required,number"`
	Address       string       `json:"address" validate:"required,lte=150"`
	Address2      string       `json:"address2" validate:"lte=150"`
	City          string       `json:"city" validate:"required,lte=100"`
	State         string       `json:"state" validate:"required,lte=100"`
	Zip           string       `json:"zip" validate:"required,number"`
	NeedPromotion bool         `json:"need_promotion"`
	Cards         []UpdateCard `json:"cards" validate:"required"`
}

type UpdateUser struct {
	UserName      string       `json:"username" validate:"required,gte=3,lte=50"`
	Name          string       `json:"name" validate:"required,lte=100"`
	Phone         string       `json:"phone" validate:"required,lte=20"`
	Address       string       `json:"address" validate:"required,lte=150"`
	Address2      string       `json:"address2" validate:"lte=150"`
	City          string       `json:"city" validate:"required,lte=100"`
	State         string       `json:"state" validate:"required,lte=100"`
	Zip           string       `json:"zip" validate:"required,lte=20"`
	NeedPromotion bool         `json:"need_promotion"`
	Cards         []UpdateCard `json:"cards" validate:"required,max=3"`
}

type AdminUpdateUser struct {
	UserName      string `json:"username" validate:"gte=3,lte=50"`
	Name          string `json:"name" validate:"lte=100"`
	Phone         string `json:"phone" validate:"lte=20"`
	Address       string `json:"address" validate:"lte=150"`
	Address2      string `json:"address2" validate:"lte=150"`
	City          string `json:"city" validate:"lte=100"`
	State         string `json:"state" validate:"lte=100"`
	Zip           string `json:"zip" validate:"lte=20"`
	NeedPromotion bool   `json:"need_promotion"`
	Email         string `json:"email" validate:"email,lte=150"`
	IsActive      bool   `json:"is_active"`
	IsAdmin       bool   `json:"is_admin"`
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

type UserDetailNoCardsListResponse = ListResponse[UserDetailNoCards]

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
