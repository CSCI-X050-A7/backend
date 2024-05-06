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
	route.Get("/:id", controller.GetMovieShows)
	routeProtectedAdmin := route.Group("", middleware.JWTProtected(), middleware.IsAdmin)
	routeProtectedAdmin.Post("/", controller.CreateMovie)
	routeProtectedAdmin.Put("/:id", controller.UpdateMovie)
	routeProtectedAdmin.Delete("/:id", controller.DeleteMovie)
}
