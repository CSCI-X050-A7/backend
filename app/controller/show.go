package controller

import (
	"github.com/gofiber/fiber/v2"
)

// Show
//
//	@Summary	get show
//	@Tags		Show
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	schema.Show
//	@Router		/api/v1/show/{id} [get]
func GetShow(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}

// Show
//
//	@Summary	create show
//	@Tags		Show
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	schema.Show
//	@Router		/api/v1/show/{id} [post]
func CreateShow(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}

// Show
//
//	@Summary	update Show
//	@Tags		Show
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	schema.Show
//	@Router		/api/v1/show/{id} [put]
func UpdateShow(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}

// Order
//
//	@Summary	delete Show
//	@Tags		Show
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	schema.Show
//	@Router		/api/v1/show/{id} [delete]
func DeleteShow(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}
