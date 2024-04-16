package controller

import (
	"github.com/gofiber/fiber/v2"
)

// Order
//
//	@Summary	get order
//	@Tags		Order
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	interface{}
//	@Router		/api/v1/order/{id} [get]
func GetOrder(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}
