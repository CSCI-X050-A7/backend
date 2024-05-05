package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/CSCI-X050-A7/backend/app/controller"
	"github.com/CSCI-X050-A7/backend/pkg/config"
	"github.com/CSCI-X050-A7/backend/pkg/middleware"
	"github.com/CSCI-X050-A7/backend/pkg/route"
	"github.com/CSCI-X050-A7/backend/platform/database"
	_ "github.com/CSCI-X050-A7/backend/platform/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// Serve ..
func Serve() {
	// connect to DB
	logrus.Infoln("Staring server...")
	database.ConnectSqlite()
	controller.NewDB(database.DB)

	app := fiber.New()

	// Attach Middlewares.
	middleware.FiberMiddleware(app)

	// Routes.
	route.GeneralRoute(app)
	route.SwaggerRoute(app)
	route.AdminRoutes(app)
	route.AuthRoutes(app)
	route.MovieRoutes(app)
	route.PromoRoutes(app)
	route.UserRoutes(app)
	route.MiscRoutes(app)
	route.CardRoutes(app)
	route.OrderRoutes(app)
	route.ShowRoutes(app)
	route.NotFoundRoute(app)

	// signal channel to capture system calls
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// start shutdown goroutine
	go func() {
		// capture sigterm and other system call here
		<-sigCh
		logrus.Infoln("Shutting down server...")
		_ = app.Shutdown()
	}()

	// start http server
	serverAddr := fmt.Sprintf("%s:%d", config.Conf.Host, config.Conf.Port)
	if err := app.Listen(serverAddr); err != nil {
		logrus.Errorf("Oops... server is not running! error: %v", err)
	}
}
