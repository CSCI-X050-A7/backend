package route

import (
	"github.com/CSCI-X050-A7/backend/app/controller"
	"github.com/CSCI-X050-A7/backend/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(a *fiber.App) {
	routeProtected := a.Group("/api/v1/users", middleware.JWTProtected())
	routeProtected.Get("/me", controller.GetUserMe)
	routeProtected.Put("/:id", controller.UpdateUser)
}
