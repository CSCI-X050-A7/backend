package controller

import (
	"fmt"
	"time"

	"github.com/CSCI-X050-A7/backend/app/schema"
	"github.com/CSCI-X050-A7/backend/pkg/config"
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
	key := []byte(config.Conf.JWTSecret)
	userDetail.CardType, _ = AESDecrypt(key, userDetail.CardType)
	userDetail.CardNumber, _ = AESDecrypt(key, userDetail.CardNumber)
	userDetail.CardExpiration, _ = AESDecrypt(key, userDetail.CardExpiration)
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
	// AES encryption for payment info
	key := []byte(config.Conf.JWTSecret)
	updateUser.CardNumber, _ = AESEncrypt(key, updateUser.CardNumber)
	updateUser.CardType, _ = AESEncrypt(key, updateUser.CardType)
	updateUser.CardExpiration, _ = AESEncrypt(key, updateUser.CardExpiration)
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

	return c.JSON(convert.To[schema.UpdateUser](user))
}

// GetOrders func get user's history orders.
//
//	@Description	a user's orders.
//	@Summary		get a user's orders
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200				{object}	schema.OrderListResponse
//	@Failure		400,401,403,404	{object}	schema.ErrorResponse	"Error"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/users/orders [get]
func GetOrders(c *fiber.Ctx) error {
	user, err := GetJWTUser(c)
	if err != nil {
		return err
	}
	logrus.Debugf("user: %v", user) // TODO: remove me
	// TODO: get orders
	return c.JSON(fiber.Map{})
}
