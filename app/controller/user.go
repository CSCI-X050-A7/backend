package controller

import (
	"fmt"
	"time"

	"github.com/CSCI-X050-A7/backend/app/model"
	"github.com/CSCI-X050-A7/backend/app/schema"
	"github.com/CSCI-X050-A7/backend/pkg/convert"
	"github.com/CSCI-X050-A7/backend/pkg/email"
	"github.com/CSCI-X050-A7/backend/pkg/validator"
	"github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
)

// GetUserMe func get a user me.
//
//	@Description	a user me.
//	@Summary		get a user me
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200				{object}	schema.UserDetail
//	@Failure		400,401,403,404	{object}	schema.ErrorResponse	"Error"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/users/me [get]
func GetUserMe(c *fiber.Ctx) error {
	user, err := GetJWTUser(c)
	if err != nil {
		return err
	}
	userDetail := convert.To[schema.UserDetail](user)
	return c.JSON(userDetail)
}

// UpdateUser update user information.
//
//	@Description	update user info.
//	@Summary		update user info
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user			body		schema.UpdateUser		true	"update user profile"
//	@Failure		400,401,403,404	{object}	schema.ErrorResponse	"Error"
//	@Success		200				{object}	schema.UserDetail		"Ok"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/users/me [put]
func UpdateUserMe(c *fiber.Ctx) error {
	user, err := GetJWTUser(c)
	if err != nil {
		return err
	}
	for _, card := range user.Cards {
		cardToDelete := card
		result := db.Delete(&cardToDelete)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"msg": result.Error,
			})
		}
	}
	updateUser := &schema.UpdateUser{}
	if err := c.BodyParser(updateUser); err != nil {
		// Return 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	// Create a new validator.
	validate := validator.NewValidator()
	if err := validate.Struct(updateUser); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":    "invalid input found",
			"errors": validator.ValidatorErrors(err),
		})
	}
	for _, card := range updateUser.Cards {
		cardToUpdate := card
		newCard := model.Card{}
		newCard.UserID = user.ID
		if err := convert.Update(&newCard, &cardToUpdate); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"msg": err.Error(),
			})
		}
		if err := db.Create(&newCard).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"msg": err.Error(),
			})
		}
	}
	updateUser.Cards = []schema.UpdateCard{}
	if err := convert.Update(&user, &updateUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	if err := db.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	go func() {
		err := email.Send(
			user.Email,
			"Cinema E-Booking System Profile Update",
			fmt.Sprintf(
				"You have updated your profile, at %s.",
				time.Now().Format(time.RFC850),
			),
		)
		if err != nil {
			logrus.Errorf("email send error: %v", err)
		}
	}()

	return c.JSON(convert.To[schema.UserDetail](user))
}

// GetUserOrders func gets the list of all orders history from the current user.
//
//	@Description	get all orders history from the current user.
//	@Summary		get all orders history
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			offset			query		integer	false	"offset"
//	@Param			limit			query		integer	false	"limit"
//	@Success		200				{object}	schema.OrderListResponse
//	@Failure		400,401,403,404	{object}	schema.ErrorResponse	"Error"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/users/me/orders [get]
func GetUserOrders(c *fiber.Ctx) error {
	// Get the current user from the authentication token
	user, err := GetJWTUser(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	pagination := GetPagination(c)
	objs, count, err := ListObjs[schema.Order](
		db.Model(model.Order{UserID: user.ID}), pagination,
	)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"offset": pagination.Offset,
		"limit":  pagination.Limit,
		"count":  count,
		"data":   objs,
	})
}
