package route

import (
	"github.com/CSCI-X050-A7/backend/app/controller"
	"github.com/CSCI-X050-A7/backend/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func OrderRoutes(a *fiber.App) {
	route := a.Group("/api/v1/orders")
	routeProtected := route.Group("", middleware.JWTProtected())
	routeProtected.Post("/", controller.CreateOrder)
	routeProtected.Get("/:id", controller.GetOrder)
	routeProtected.Post("/:id/checkout", controller.CheckoutOrder)
}
