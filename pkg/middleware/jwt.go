package middleware

import (
	config "github.com/CSCI-X050-A7/backend/pkg/config"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// JWTProtected func for specify route group with JWT authentication.
func JWTProtected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: []byte(config.Conf.JWTSecret)},
		TokenLookup: "header:Authorization,cookie:access_token",
	})
}
