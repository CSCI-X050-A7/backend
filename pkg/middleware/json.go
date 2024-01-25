package middleware

import "github.com/gofiber/fiber/v2"

func Json(c *fiber.Ctx) error {
	c.Accepts(fiber.MIMEApplicationJSON)
	return c.Next()
}
