package controller

import (
	"fmt"

	"github.com/CSCI-X050-A7/backend/app/model"
	"github.com/CSCI-X050-A7/backend/app/schema"
	"github.com/CSCI-X050-A7/backend/pkg/convert"
	"github.com/CSCI-X050-A7/backend/pkg/email"
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
//	@Security		ApiKeyAuth
//	@Router			/api/v1/orders/{id} [get]
func GetOrder(c *fiber.Ctx) error {
	user, err := GetJWTUser(c)
	if err != nil {
		return err
	}
	ID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	order := model.Order{ID: ID}
	err = db.Preload("Tickets").First(&order).Error
	if err != nil || user.ID != order.UserID {
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
//	@Param			Order		body		schema.CreateOrder		true	"Create new order"
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
	createOrder := &schema.CreateOrder{}
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

	// Look up the promotion code in the database
	promotion := model.Promotion{}
	if createOrder.PromotionCode != "" {
		err = db.Where(&model.Promotion{Code: createOrder.PromotionCode}).
			First(&promotion).Error
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"msg": "promotion not found",
			})
		}
	}

	// Look up the show in the database
	show := model.Show{}
	err = db.Where(&model.Show{ID: createOrder.ShowID}).
		First(&show).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "show not found",
		})
	}
	movie := model.Movie{}
	err = db.Where(&model.Movie{ID: show.MovieID}).
		First(&movie).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "movie not found",
		})
	}

	newOrder := model.Order{
		ShowID:      show.ID,
		UserID:      user.ID,
		PromotionID: promotion.ID,
		CheckOut:    false,
		MovieTitle:  movie.Title,
	}
	if err := db.Create(&newOrder).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	// create tickets
	var totalPrice float64
	totalPrice = 0
	for _, ticket := range createOrder.Tickets {
		price := show.AdultTicketPrice
		if ticket.Type != "child" && ticket.Type != " senior" {
			ticket.Type = "adult"
		}
		if ticket.Type == "child" {
			price = show.ChildTicketPrice
		} else if ticket.Type == "senior" {
			price = show.SeniorTicketPrice
		}
		newTicket := model.Ticket{
			OrderID: newOrder.ID,
			ShowID:  show.ID,
			Seat:    ticket.Seat,
			Type:    ticket.Type,
			Price:   price,
		}
		totalPrice += price
		if err := db.Create(&newTicket).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"msg": err.Error(),
			})
		}
	}
	newOrder.TicketPrice = totalPrice

	newOrder.BookingFeePrice = show.BookingFee
	totalPrice += show.BookingFee

	if createOrder.PromotionCode != "" {
		promotionPrice := -totalPrice * promotion.Discount
		newOrder.PromotionPrice = promotionPrice
		totalPrice += promotionPrice
	}

	// sales tax
	salesTaxPrice := totalPrice * 0.05
	newOrder.SalesTaxPrice = salesTaxPrice
	totalPrice += salesTaxPrice

	newOrder.TotalPrice = totalPrice

	if err := db.Save(&newOrder).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	err = db.Preload("Tickets").Where(&model.Order{ID: newOrder.ID}).
		First(&newOrder).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.JSON(convert.To[schema.Order](newOrder))
}

// CheckoutOrder func for check out an existing order.
//
//	@Description	Checkout a order.
//	@Summary		checkout a order
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string					true	"Order ID"
//	@Param			Card		body		schema.UpdateCard		true	"Card to checkout"
//	@Failure		400,401,500	{object}	schema.ErrorResponse	"Error"
//	@Success		200			{object}	schema.Order			"Ok"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/orders/{id}/checkout [post]
func CheckoutOrder(c *fiber.Ctx) error {
	user, err := GetJWTUser(c)
	if err != nil {
		return err
	}
	ID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	order := model.Order{ID: ID}
	err = db.First(&order).Error
	if err != nil || user.ID != order.UserID {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "order not found",
		})
	}
	if order.CheckOut {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "order checked out",
		})
	}

	// Create new card struct
	card := &schema.UpdateCard{}
	if err := c.BodyParser(card); err != nil {
		// Return 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	// Create a new validator for a Order model.
	validate := validator.NewValidator()
	if err := validate.Struct(card); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":    "invalid input found",
			"errors": validator.ValidatorErrors(err),
		})
	}
	newCard := model.Card{}
	newCard.UserID = user.ID
	if err := convert.Update(&newCard, &card); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	if err := db.Create(&newCard).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	order.CardID = newCard.ID
	order.CheckOut = true

	if err := db.Save(&order).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	go func() {
		err := email.Send(
			user.Email,
			"Cinema E-Booking System Checkout Confirmation",
			fmt.Sprintf("You have checkout a order with ID %s, movie title %s, total price %f",
				order.ID, order.MovieTitle, order.TotalPrice),
		)
		if err != nil {
			logrus.Errorf("email send error: %v", err)
		}
	}()

	err = db.Preload("Tickets").Where(&model.Order{ID: order.ID}).
		First(&order).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.JSON(convert.To[schema.Order](order))
}
