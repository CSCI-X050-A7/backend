package route

import (
	"github.com/CSCI-X050-A7/backend/app/controller"
	"github.com/CSCI-X050-A7/backend/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func MovieRoutes(a *fiber.App) {
	route := a.Group("/api/v1/movies")
	route.Get("/", controller.GetMovies)
	route.Get("/:id", controller.GetMovie)
	routeProtected := route.Group("", middleware.JWTProtected())
	routeProtected.Post("/", controller.CreateMovie)
	routeProtected.Put("/:id", controller.UpdateMovie)
	routeProtected.Delete("/:id", controller.DeleteMovie)
}
