package controller

import (
	"fmt"

	"github.com/CSCI-X050-A7/backend/app/model"
	"github.com/CSCI-X050-A7/backend/app/schema"
	"github.com/CSCI-X050-A7/backend/pkg/convert"
	"github.com/CSCI-X050-A7/backend/pkg/email"
	"github.com/CSCI-X050-A7/backend/pkg/validator"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
)

// GetPromos func gets all promotions.
//
// @Description	Get all promotions.
// @Summary		get all promotions
// @Tags		Admin
// @Accept		json
// @Produce		json
// @Param       search query    string  false "search by code"
// @Param 		offset query    integer false "offset"
// @Param 		limit  query    integer false "limit"
// @Success		200	   {object}	schema.PromoListResponse
// @Failure		400	   {object}	schema.ErrorResponse	"Error"
// @Router		/api/v1/admin/promotions [get]
func GetPromos(c *fiber.Ctx) error {
	pagination := GetPagination(c)
	search := c.Query("search", "")
	statement := db.Model(model.Promotion{})
	if search != "" {
		statement = statement.
			Where("LOWER(code) LIKE ?", fmt.Sprintf("%%%s%%", search))
	}
	objs, count, err := ListObjs[schema.Promotion](
		statement,
		pagination,
	)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"offset": pagination.Offset,
		"limit":  pagination.Limit,
		"data":   objs,
		"count":  count,
	})
}

// CreatePromo func creates a new promotion
//
// @Description	Create a new promotion.
// @Summary		create a new promotion
// @Tags		Admin
// @Accept		json
// @Produce		json
// @Param       promotion body schema.UpsertPromotion true "Create new promo"
// @Success		200	   {object}	schema.Promotion
// @Failure		400	   {object}	schema.ErrorResponse	"Error"
// @Router		/api/v1/admin/promotions [post]
// @Security    ApiKeyAuth
func CreatePromo(c *fiber.Ctx) error {
	createPromo := &schema.UpsertPromotion{}
	if err := c.BodyParser(createPromo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	logrus.Infof("CreatePromo: %+v", createPromo)

	// TODO: add validator

	newPromo := model.Promotion{}
	if err := convert.Update(&newPromo, &createPromo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	if err := db.Create(&newPromo).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	var promoUsers []model.User
	db.Where(&model.User{NeedPromotion: true}).Find(&promoUsers)

	for _, sendUser := range promoUsers {
		go func(u model.User) {
			err := email.Send(
				u.Email,
				"Don't miss out on our new promotional offer!",
				fmt.Sprintf("Use promo code '%s' at checkout for a %.0f%% off discount!.",
					newPromo.Code, newPromo.Discount*100),
			)
			if err != nil {
				logrus.Errorf("email send error: %v", err)
			}
		}(sendUser)
	}
	return c.JSON(convert.To[schema.Promotion](newPromo))
}

// UpdatePromo func update a promo.
//
//	@Description	update promo
//	@Summary		update a promo
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			id					path		string			true	"Promo ID"
//	@Param			updatepromo			body		schema.UpsertPromotion	true	"Update a promo"
//	@Success		200					{object}	schema.Promotion
//	@Failure		400,401,403,404,500	{object}	schema.ErrorResponse	"Error"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/admin/promotions/{id} [put]
func UpdatePromo(c *fiber.Ctx) error {
	ID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	promotion := model.Promotion{ID: ID}
	err = db.First(&promotion).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "promo not found",
		})
	}

	updatePromo := &schema.UpsertPromotion{}
	if err := c.BodyParser(updatePromo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	validate := validator.NewValidator()
	if err := validate.Struct(updatePromo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":    "invalid input found",
			"errors": validator.ValidatorErrors(err),
		})
	}

	if err := convert.Update(&promotion, &updatePromo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	if err := db.Save(&promotion).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	return c.JSON(convert.To[schema.Promotion](promotion))
}
