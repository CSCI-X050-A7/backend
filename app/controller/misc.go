package controller

import (
	"github.com/CSCI-X050-A7/backend/pkg/config"

	"github.com/gofiber/fiber/v2"
)

// Version
//
//	@Summary	get current software version
//	@Tags		Misc
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	interface{}
//	@Router		/api/v1/version [get]
func Version(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"GoVersion": config.GoVersion,
		"GitAuthor": config.GitAuthor,
		"GitCommit": config.GitCommit,
		"BuiltAt":   config.BuiltAt,
		"Version":   config.Version,
	})
}
