package main

import (
	"github.com/CSCI-X050-A7/backend/cmd/server"
	_ "github.com/CSCI-X050-A7/backend/docs" // load API Docs files (Swagger)
	_ "github.com/CSCI-X050-A7/backend/pkg/config"
)

// @title Fiber Go API
// @version 1.0
// @description Fiber go web framework based REST API boilerplate
// @termsOfService
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @host localhost:5000
// @BasePath /api
func main() {
	// setup various configuration for app
	server.Serve()
}
