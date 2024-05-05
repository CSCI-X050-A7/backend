package route

import (
	"github.com/CSCI-X050-A7/backend/app/controller"

	"github.com/gofiber/fiber/v2"
)

func OrderRoutes(a *fiber.App) {
	route := a.Group("/api/v1/orders")
	route.Get("/:id", controller.GetOrder)
	route.Post("/:id", controller.CreateOrder)
	route.Put("/:id", controller.UpdateOrder)
	route.Delete("/:id", controller.DeleteOrder)
}
