package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v5"
)

func FiberMiddleware(a *fiber.App) {
	a.Use(
		// Add CORS to each route.
		cors.New(),
		// Add simple logger.
		logger.New(logger.Config{
			Format:     "${time} ${magenta}[${ip}:${port}]${reset} ${status} - ${yellow}[${latency}]${reset} ${method} ${path}\n",
			TimeFormat: "2006-01-02 15:04:03.000",
		}),
		Json,
	)
}

func IsAdmin(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	isAdmin, ok := claims["admin"]
	if !ok || !isAdmin.(bool) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"msg": "Forbidden",
		})
	}

	return c.Next()
}
