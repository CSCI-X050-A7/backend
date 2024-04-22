package route

import (
	"github.com/CSCI-X050-A7/backend/app/controller"

	"github.com/gofiber/fiber/v2"
)

func ShowRoutes(a *fiber.App) {
	route := a.Group("/api/v1/shows")
	route.Get("/:id", controller.GetShow)
	route.Post("/:id", controller.CreateShow)
	route.Put("/:id", controller.UpdateShow)
	route.Delete("/:id", controller.DeleteShow)
}
