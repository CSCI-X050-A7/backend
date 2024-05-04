package route

import (
	"github.com/CSCI-X050-A7/backend/app/controller"
	"github.com/CSCI-X050-A7/backend/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func PromoRoutes(a *fiber.App) {
	route := a.Group("/api/v1/promotions")
	route.Get("/", controller.GetPromos)
	route.Get("/:id", controller.GetPromo)
	routeProtectedAdmin := route.Group("", middleware.JWTProtected(), middleware.IsAdmin)
	routeProtectedAdmin.Post("/", controller.CreatePromo)
	routeProtectedAdmin.Put("/:id", controller.UpdatePromo)
	routeProtectedAdmin.Delete("/:id", controller.DeletePromo)
}
