package route

import (
	"github.com/CSCI-X050-A7/backend/app/controller"
	"github.com/CSCI-X050-A7/backend/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func TicketRoutes(a *fiber.App) {
	route := a.Group("/api/v1/tickets")
	route.Get("/:id", controller.GetTicket)
	routeProtectedAdmin := route.Group("", middleware.JWTProtected(), middleware.IsAdmin)
	routeProtectedAdmin.Post("/", controller.CreateTicket)
	routeProtectedAdmin.Put("/:id", controller.UpdateTicket)
	// routeProtectedAdmin.Delete("/:id", controller.DeleteTicket)
}
