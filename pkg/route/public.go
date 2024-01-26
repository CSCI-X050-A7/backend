package route

import (
	"github.com/CSCI-X050-A7/backend/app/controller"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public route.
func PublicRoutes(a *fiber.App) {
	// Create route group.
	route := a.Group("/api/v1")

	route.Post("/auth/login", controller.Login)
	route.Get("/books", controller.GetBooks)
	route.Get("/books/:id", controller.GetBook)
}
