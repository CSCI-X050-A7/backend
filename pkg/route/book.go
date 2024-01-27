package route

import (
	"github.com/CSCI-X050-A7/backend/app/controller"
	"github.com/CSCI-X050-A7/backend/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func BookRoutes(a *fiber.App) {
	route := a.Group("/api/v1/books")
	route.Get("/", controller.GetBooks)
	route.Get("/:id", controller.GetBook)
	routeProtected := route.Group("", middleware.JWTProtected())
	routeProtected.Post("/", controller.CreateBook)
	routeProtected.Put("/:id", controller.UpdateBook)
	routeProtected.Delete("/:id", controller.DeleteBook)
}
