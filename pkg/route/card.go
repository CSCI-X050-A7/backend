package route

import (
	"github.com/CSCI-X050-A7/backend/app/controller"

	"github.com/gofiber/fiber/v2"
)

func CardRoutes(a *fiber.App) {
	route := a.Group("/api/v1/cards")
	route.Get("/:id", controller.GetCard)
	// routeProtectedAdmin := route.Group("", middleware.JWTProtected(), middleware.IsAdmin)
}
