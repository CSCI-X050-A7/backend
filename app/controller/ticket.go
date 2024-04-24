package controller

import (
	"github.com/CSCI-X050-A7/backend/app/model"
	"github.com/CSCI-X050-A7/backend/app/schema"
	"github.com/CSCI-X050-A7/backend/pkg/convert"
	"github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
)

// CreateTicket func creates a new ticket
//
// @Description	Create a new ticket.
// @Summary		create a new ticket
// @Tags		Ticket
// @Accept		json
// @Produce		json
// @Param       ticket body schema.UpsertTicket true "Create new ticket"
// @Success		200	   {object}	schema.Ticket
// @Failure		400	   {object}	schema.ErrorResponse	"Error"
// @Router		/api/v1/tickets [post]
// @Security    ApiKeyAuth
func CreateTicket(c *fiber.Ctx) error {

	createTicket := &schema.Ticket{}
	if err := c.BodyParser(createTicket); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	logrus.Infof("CreateTicket: %+v", createTicket)

	newTicket := model.Ticket{}
	if err := convert.Update(&newTicket, &createTicket); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	if err := db.Create(&newTicket).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	return c.JSON(convert.To[schema.Ticket](newTicket))
}

// UpdateTicket func updates a ticket
//
// @Description	Update a ticket.
// @Summary		update a ticket
// @Tags		Ticket
// @Param       id path string true "Ticket ID"
// @Param       ticket body schema.UpsertTicket true "Update ticket"
// @Success		200	   {object}	schema.Ticket
// @Failure		400	   {object}	schema.ErrorResponse	"Error"
// @Router		/api/v1/tickets/{id} [put]
// @Security    ApiKeyAuth
func UpdateTicket(c *fiber.Ctx) error {
	id := c.Params("id")

	// Check if the ticket exists
	existingTicket := model.Ticket{}
	if err := db.First(&existingTicket, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "Ticket not found",
		})
	}

	updateTicket := &schema.Ticket{}
	if err := c.BodyParser(updateTicket); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	logrus.Infof("UpdateTicket: %+v", updateTicket)

	if err := convert.Update(&existingTicket, &updateTicket); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	if err := db.Save(&existingTicket).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.JSON(convert.To[schema.Ticket](existingTicket))
}

// GetTicket func retrieves a ticket by ID
//
// @Description	Retrieve a ticket by ID.
// @Summary		retrieve a ticket
// @Tags		Ticket
// @Param       id path string true "Ticket ID"
// @Success		200	   {object}	schema.Ticket
// @Failure		400	   {object}	schema.ErrorResponse	"Error"
// @Router		/api/v1/tickets/{id} [get]
// @Security    ApiKeyAuth
func GetTicket(c *fiber.Ctx) error {
	id := c.Params("id")

	// Check if the ticket exists
	existingTicket := model.Ticket{}
	if err := db.First(&existingTicket, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "Ticket not found",
		})
	}

	return c.JSON(convert.To[schema.Ticket](existingTicket))
}
