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
//	@Success	200	{object}	schema.Order
//	@Router		/api/v1/order/{id} [get]
func GetOrder(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}

// Order
//
//	@Summary	place order
//	@Tags		Order
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	schema.Order
//	@Router		/api/v1/order/{id} [post]
func PlaceOrder(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}

// Order
//
//	@Summary	update order
//	@Tags		Order
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	schema.Order
//	@Router		/api/v1/order/{id} [put]
func UpdateOrder(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}

// Order
//
//	@Summary	delete order
//	@Tags		Order
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	schema.Order
//	@Router		/api/v1/order/{id} [delete]
func DeleteOrder(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}
