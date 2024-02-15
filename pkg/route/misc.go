package route

import (
	"github.com/CSCI-X050-A7/backend/app/controller"

	"github.com/gofiber/fiber/v2"
)

func MiscRoutes(a *fiber.App) {
	route := a.Group("/api/v1")
	route.Get("/version", controller.Version)
}
