package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func GeneralRoute(a *fiber.App) {
	a.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg":  "Welcome to Fiber Go API!",
			"docs": "/swagger/index.html",
		})
	})
}

func SwaggerRoute(a *fiber.App) {
	// Create route group.
	route := a.Group("/swagger")
	route.Get("*", swagger.HandlerDefault)
}

func NotFoundRoute(a *fiber.App) {
	a.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"msg": "sorry, endpoint is not found",
			})
		},
	)
}
