package route

import (
	"github.com/CSCI-X050-A7/backend/app/controller"
	"github.com/CSCI-X050-A7/backend/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(a *fiber.App) {
	route := a.Group("/api/v1/auth")
	route.Post("/login", controller.Login)
	route.Post("/register", controller.Register)
	route.Post("/activate", controller.Activate)
	route.Post("/forgotpassword", controller.ForgotPassword) // forgot password
	route.Post("/resetpassword", controller.ResetPassword)   // reset password
	routeProtected := route.Group("", middleware.JWTProtected())
	routeProtected.Post("/logout", controller.Logout)
	routeProtected.Post("/jwt", controller.JWT)
}
