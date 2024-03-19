package model

import (
	"math/rand"
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

// generates a random string of specified length
func GenerateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomString := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := range randomString {
		randomString[i] = charset[rand.Intn(len(charset))]
	}
	return string(randomString), nil
}

// UpdatePasswordCode generates a random string and sets it to the PasswordCode field
func (u *User) UpdatePasswordCode() error {
	// Generate a random string
	randomString, err := GenerateRandomString(10)
	if err != nil {
		return err
	}
	u.PasswordCode = randomString
	return nil
}
