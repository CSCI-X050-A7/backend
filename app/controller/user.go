package controller

import (
	"github.com/CSCI-X050-A7/backend/app/model"
	"github.com/CSCI-X050-A7/backend/app/schema"
	"github.com/CSCI-X050-A7/backend/pkg/convert"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

// GetUserMe func get a user me.
// @Description a user me.
// @Summary get a user me
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} schema.User
// @Failure 400,401,403,404 {object} schema.ErrorResponse "Error"
// @Security ApiKeyAuth
// @Router /api/v1/users/me [get]
func GetUserMe(c *fiber.Ctx) error {
	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	ID, err := uuid.Parse(claims["user_id"].(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	user := model.User{ID: ID}
	err = db.First(&user).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "user not found",
		})
	}
	return c.JSON(convert.To[schema.User](user))
}
