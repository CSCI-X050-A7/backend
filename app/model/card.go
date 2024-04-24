package model

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/CSCI-X050-A7/backend/pkg/config"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Card struct to describe Card object.
type Card struct {
	ID         uuid.UUID `gorm:"primarykey;type:uuid;default:(uuid_generate_v4())"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Type       string
	Number     string
	Expiration string
	Address    string
	Address2   string
	City       string
	State      string
	Zip        string
}

func AESEncrypt(key []byte, plaintext string) (string, error) {
	truncateKey := make([]byte, 16)
	copy(truncateKey, key)
	var byteString []byte
	if len(plaintext) < 16 {
		paddedTxt := make([]byte, 16)
		copy(paddedTxt, plaintext)
		byteString = []byte(paddedTxt)
	} else {
		byteString = []byte(plaintext)
	}
	block, err := aes.NewCipher(truncateKey)
	// err: cipher cannot be created
	if err != nil {
		return "", err
	}
	cipherText := make([]byte, aes.BlockSize+len(byteString))
	iv := cipherText[:aes.BlockSize] // extends to blocksize
	// err: can't encrypt
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], byteString)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func AESDecrypt(key []byte, plaintext string) (string, error) {
	cipherText, err := base64.StdEncoding.DecodeString(plaintext)
	// err: cannot base64 decode
	if err != nil {
		return "", err
	}

	truncateKey := make([]byte, 16)
	copy(truncateKey, key)
	block, err := aes.NewCipher(truncateKey)
	// err: cannot make new cipher
	if err != nil {
		return "", err
	}
	// err: ciphertext block size is wrong
	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext wrong block size")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)
	decrypted := string(cipherText)
	decrypted = strings.Trim(decrypted, "\u0000")
	return decrypted, nil
}

func (c *Card) BeforeSave(tx *gorm.DB) (err error) {
	key := []byte(config.Conf.JWTSecret)
	c.Number, _ = AESEncrypt(key, c.Number)
	c.Type, _ = AESEncrypt(key, c.Type)
	c.Expiration, _ = AESEncrypt(key, c.Expiration)
	return
}

func (c *Card) AfterFind(tx *gorm.DB) (err error) {
	key := []byte(config.Conf.JWTSecret)
	c.Number, _ = AESDecrypt(key, c.Number)
	c.Type, _ = AESDecrypt(key, c.Type)
	c.Expiration, _ = AESDecrypt(key, c.Expiration)
	return
}
