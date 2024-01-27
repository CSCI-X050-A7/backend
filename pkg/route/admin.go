package route

import (
	"github.com/CSCI-X050-A7/backend/app/controller"
	"github.com/CSCI-X050-A7/backend/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(a *fiber.App) {
	routeProtectedAdmin := a.Group("/api/v1/admin", middleware.JWTProtected(), middleware.IsAdmin)
	routeProtectedAdmin.Post("/users", controller.AdminCreateUser)
	routeProtectedAdmin.Get("/users", controller.AdminGetUsers)
	routeProtectedAdmin.Get("/users/:id", controller.AdminGetUser)
	routeProtectedAdmin.Put("/users/:id", controller.AdminUpdateUser)
	routeProtectedAdmin.Delete("/users/:id", controller.AdminDeleteUser)
}
