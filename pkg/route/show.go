package route

import (
	"github.com/CSCI-X050-A7/backend/app/controller"
	"github.com/CSCI-X050-A7/backend/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func ShowRoutes(a *fiber.App) {
	route := a.Group("/api/v1/shows")
	route.Get("/:id", controller.GetShow)
	routeProtectedAdmin := route.Group("", middleware.JWTProtected(), middleware.IsAdmin)
	routeProtectedAdmin.Post("/:id", controller.CreateShow)
	routeProtectedAdmin.Put("/:id", controller.UpdateShow)
	routeProtectedAdmin.Delete("/:id", controller.DeleteShow)
}
