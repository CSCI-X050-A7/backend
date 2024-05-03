package route

import (
	"github.com/CSCI-X050-A7/backend/app/controller"
	"github.com/CSCI-X050-A7/backend/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func CardRoutes(a *fiber.App) {
	route := a.Group("/api/v1/cards")
	routeProtected := route.Group("", middleware.JWTProtected())
	routeProtected.Get("/:id", controller.GetCard)
}
