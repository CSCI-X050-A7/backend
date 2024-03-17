package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct to describe User object.
type User struct {
	ID             uuid.UUID `gorm:"primarykey;type:uuid;default:(uuid_generate_v4())"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	IsActive       bool
	ActivationCode string
	IsAdmin        bool
	UserName       string
	Email          string
	Password       string
	PasswordCode   string
	Name           string
	Phone          string
	Address        string
	Address2       string
	City           string
	State          string
	Zip            string
	CardType       string
	CardNumber     string
	CardExpiration string
	CardAddress    string
	CardAddress2   string
	CardCity       string
	CardState      string
	CardZip        string
	NeedPromotion  bool
}
