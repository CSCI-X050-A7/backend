package controller

import (
	"fmt"

	"github.com/CSCI-X050-A7/backend/app/model"
	"github.com/CSCI-X050-A7/backend/app/schema"
	"github.com/CSCI-X050-A7/backend/pkg/convert"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
)

// GetPromos func gets all promotions.
//
// @Description	Get all promotions.
// @Summary		get all promotions
// @Tags		Promotion
// @Accept		json
// @Produce		json
// @Param       search query    string  false "search by title"
// @Param       movie  query    string  false "search by movie"
// @Param 		offset query    integer false "offset"
// @Param 		limit  query    integer false "limit"
// @Success		200	   {object}	schema.PromoListResponse
// @Failure		400	   {object}	schema.ErrorResponse	"Error"
// @Router		/api/v1/promotions [get]
func GetPromos(c *fiber.Ctx) error {
	pagination := GetPagination(c)
	search := c.Query("search", "")
	movie := c.Query("movie", "")
	statement := db.Model(model.Promotion{})
	if search != "" {
		statement = statement.
			Where("LOWER(title) LIKE ?", fmt.Sprintf("%%%s%%", search))
	}
	if movie != "" {
		statement = statement.
			Where("LOWER(movie_affected) LIKE ?", movie)
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
// @Tags		Promotion
// @Accept		json
// @Produce		json
// @Param       promotion body schema.UpsertPromotion true "Create new promo"
// @Success		200	   {object}	schema.Promotion
// @Failure		400	   {object}	schema.ErrorResponse	"Error"
// @Router		/api/v1/promotions [post]
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
	return c.JSON(convert.To[schema.Promotion](newPromo))
}

// UpdatePromo func update a promo.
//
//	@Description	update promo
//	@Summary		update a promo
//	@Tags			Promotion
//	@Accept			json
//	@Produce		json
//	@Param			id					path		string			true	"Promo ID"
//	@Param			updatepromo			body		schema.UpsertPromotion	true	"Update a promo"
//	@Success		200					{object}	schema.Promotion
//	@Failure		400,401,403,404,500	{object}	schema.ErrorResponse	"Error"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/promotions/{id} [put]
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

	// TODO: validator

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

// DeletePromo func delete a promotion
// @Description     delete promotion
// @Summary 	    delete a promotion
// @Tags            Promotion
// @Accept          json
// @Produce         json
// @Param           id                 path        string                true      "Promotion ID"
// @Success         200                {object}    schema.Promotion      "Ok"
// @Failure         500,401,403,404    {object}    schema.ErrorResponse  "Error"
// @Security        ApiKeyAuth
// @Router          /api/v1/promotions/{id} [delete]
func DeletePromo(c *fiber.Ctx) error {
	ID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	promotion := model.Promotion{ID: ID}
	err = db.First(&promotion).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "promotion not found",
		})
	}

	promotion = model.Promotion{ID: ID}
	result := db.Delete(&promotion)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": result.Error,
		})
	}
	return c.JSON(fiber.Map{})
}

// GetPromo func get a promotion
// @Description     a promotion
// @Summary 	    get a promotion
// @Tags            Promotion
// @Accept          json
// @Produce         json
// @Param           id                 path        string                true      "Promotion ID"
// @Success         200                {object}    schema.Promotion      "Ok"
// @Failure         400,404            {object}    schema.ErrorResponse  "Error"
// @Router 		    /api/v1/promotions/{id} [get]
func GetPromo(c *fiber.Ctx) error {
	ID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	promotion := model.Promotion{ID: ID}
	err = db.First(&promotion).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "promotion not found",
		})
	}
	return c.JSON(convert.To[schema.Promotion](promotion))
}
