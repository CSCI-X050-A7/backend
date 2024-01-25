package route

import (
	"github.com/CSCI-X050-A7/backend/app/controller"
	"github.com/CSCI-X050-A7/backend/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private route.
func PrivateRoutes(a *fiber.App) {
	// Admin route group
	adminRoute := a.Group("/api/v1/users", middleware.JWTProtected(), middleware.IsAdmin)
	// User
	adminRoute.Post("/", controller.CreateUser)
	adminRoute.Get("/", controller.GetUsers)
	adminRoute.Get("/:id", controller.GetUser)
	adminRoute.Put("/:id", controller.UpdateUser)
	adminRoute.Delete("/:id", controller.DeleteUser)

	// Book
	route := a.Group("/api/v1/books", middleware.JWTProtected())
	route.Post("/", controller.CreateBook)
	route.Put("/:id", controller.UpdateBook)
	route.Delete("/:id", controller.DeleteBook)
}
