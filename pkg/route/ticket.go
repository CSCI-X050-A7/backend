package route

import (
	"github.com/CSCI-X050-A7/backend/app/controller"
	"github.com/gofiber/fiber/v2"
)

func TicketRoutes(a *fiber.App) {
	route := a.Group("/api/v1/tickets")
	route.Get("/:id", controller.GetTicket)
}
