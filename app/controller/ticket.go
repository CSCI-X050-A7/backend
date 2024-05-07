package controller

import (
	"github.com/CSCI-X050-A7/backend/app/model"
	"github.com/CSCI-X050-A7/backend/app/schema"
	"github.com/CSCI-X050-A7/backend/pkg/convert"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

// GetTicket func retrieves a ticket by ID
//
// @Description	Retrieve a ticket by ID.
// @Summary		retrieve a ticket by ID
// @Tags		Ticket
// @Accept		json
// @Produce		json
// @Param		id		path		string	true	"Ticket ID"
// @Success		200		{object}	schema.Ticket
// @Failure		400,404	{object}	schema.ErrorResponse	"Error"
// @Router		/api/v1/tickets/{id} [get]
// @Security    ApiKeyAuth
func GetTicket(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "Invalid ticket ID",
		})
	}
	// Check if the ticket exists
	existingTicket := model.Ticket{}
	if err := db.Where("id = ?", id).First(&existingTicket).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "Ticket not found",
		})
	}

	return c.JSON(convert.To[schema.Ticket](existingTicket))
}
