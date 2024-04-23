package controller

import (
	"github.com/CSCI-X050-A7/backend/app/model"
	"github.com/CSCI-X050-A7/backend/app/schema"
	"github.com/CSCI-X050-A7/backend/pkg/convert"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetCard(c *fiber.Ctx) error {
	ID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	card := model.Card{ID: ID}
	err = db.First(&card).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "card not found",
		})
	}
	return c.JSON(convert.To[schema.Card](card))
}
