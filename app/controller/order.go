package controller

import (
	"github.com/CSCI-X050-A7/backend/app/model"
	"github.com/CSCI-X050-A7/backend/app/schema"
	"github.com/CSCI-X050-A7/backend/pkg/convert"
	"github.com/CSCI-X050-A7/backend/pkg/validator"
	"github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetOrder func gets a order.
//
//	@Description	a order.
//	@Summary		get a order
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"Order ID"
//	@Success		200		{object}	schema.Order
//	@Failure		400,404	{object}	schema.ErrorResponse	"Error"
//	@Router			/api/v1/orders/{id} [get]
func GetOrder(c *fiber.Ctx) error {
	ID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	order := model.Order{ID: ID}
	err = db.First(&order).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "Order not found",
		})
	}
	return c.JSON(convert.To[schema.Order](order))
}

// CreateOrder func for creates a new order.
//
//	@Description	Create a new order.
//	@Summary		create a new order
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			Order		body		schema.UpsertOrder		true	"Create new order"
//	@Failure		400,401,500	{object}	schema.ErrorResponse	"Error"
//	@Success		200			{object}	schema.Order			"Ok"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/orders [post]
func CreateOrder(c *fiber.Ctx) error {
	user, err := GetJWTUser(c)
	if err != nil {
		return err
	}
	// Create new order struct
	createOrder := &schema.UpsertOrder{}
	if err := c.BodyParser(createOrder); err != nil {
		// Return 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	logrus.Infof("create order: %+v", createOrder)

	// Create a new validator for a Order model.
	validate := validator.NewValidator()
	if err := validate.Struct(createOrder); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":    "invalid input found",
			"errors": validator.ValidatorErrors(err),
		})
	}

	// Look up the promotion in the database
	var promotion model.Promotion
	db.Where("code = ?", createOrder.Promotion).First(&promotion)

	// Look up the tickets in the database
	var tickets []model.Ticket
	db.Where("id IN ?", createOrder.TicketsArray).Find(&tickets)

	// Look up the show in the database
	var show model.Show
	db.Where("id = ?", createOrder.Show).First(&show)

	// Look up the card in the database
	var card model.Card
	db.Where("id = ?", createOrder.Card).First(&card)

	newOrder := model.Order{
		ShowID:      show.ID,
		CardID:      card.ID,
		UserID:      user.ID,
		PromotionID: promotion.ID,
		Tickets:     tickets,
	}
	if err := convert.Update(&newOrder, &createOrder); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	if err := db.Create(&newOrder).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.JSON(convert.To[schema.Order](newOrder))
}

// UpdateOrder func update a Order.
//
//	@Description	update order
//	@Summary		update a order
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			id					path		string			true	"Order ID"
//	@Param			updateOrder			body		schema.UpsertOrder	true	"Update a Order"
//	@Success		200					{object}	schema.Order
//	@Failure		400,401,403,404,500	{object}	schema.ErrorResponse	"Error"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/orders/{id} [put]
func UpdateOrder(c *fiber.Ctx) error {
	ID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	order := model.Order{ID: ID}
	err = db.First(&order).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "Order not found",
		})
	}

	updateOrder := &schema.UpsertOrder{}
	if err := c.BodyParser(updateOrder); err != nil {
		// Return 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	// Create a new validator for a order model.
	validate := validator.NewValidator()
	if err := validate.Struct(updateOrder); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":    "invalid input found",
			"errors": validator.ValidatorErrors(err),
		})
	}

	if err := convert.Update(&order, &updateOrder); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	if err := db.Save(&order).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.JSON(convert.To[schema.Order](order))
}

// DeleteOrder func delete a Order.
//
//	@Description	delete order
//	@Summary		delete a order
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string					true	"Order ID"
//	@Success		200				{object}	interface{}				"Ok"
//	@Failure		401,403,404,500	{object}	schema.ErrorResponse	"Error"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/orders/{id} [delete]
func DeleteOrder(c *fiber.Ctx) error {
	ID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	order := model.Order{ID: ID}
	err = db.First(&order).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "user not found",
		})
	}

	order = model.Order{ID: ID}
	result := db.Delete(&order)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": result.Error,
		})
	}

	return c.JSON(fiber.Map{})
}
