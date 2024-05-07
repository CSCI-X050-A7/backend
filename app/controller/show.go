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

// GetShow func gets a Show.
//
//	@Description	a show.
//	@Summary		get a show
//	@Tags			Show
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"Show ID"
//	@Success		200		{object}	schema.Show
//	@Failure		400,404	{object}	schema.ErrorResponse	"Error"
//	@Router			/api/v1/Shows/{id} [get]
func GetShow(c *fiber.Ctx) error {
	ID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	show := model.Show{ID: ID}
	err = db.First(&show).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "Show not found",
		})
	}
	return c.JSON(convert.To[schema.Show](show))
}

// CreateShow func for creates a new show.
//
//	@Description	Create a new show.
//	@Summary		create a new show
//	@Tags			Show
//	@Accept			json
//	@Produce		json
//	@Param			Show		body		schema.UpsertShow		true	"Create new Show"
//	@Failure		400,401,500	{object}	schema.ErrorResponse	"Error"
//	@Success		200			{object}	schema.Show			"Ok"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/Shows [post]
func CreateShow(c *fiber.Ctx) error {
	// Create new Show struct
	createShow := &schema.UpsertShow{}
	if err := c.BodyParser(createShow); err != nil {
		// Return 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	logrus.Infof("create Show: %+v", createShow)

	// Create a new validator for a Show model.
	validate := validator.NewValidator()
	if err := validate.Struct(createShow); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":    "invalid input found",
			"errors": validator.ValidatorErrors(err),
		})
	}

	newShow := model.Show{}
	if err := convert.Update(&newShow, &createShow); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	if err := db.Create(&newShow).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.JSON(convert.To[schema.Show](newShow))
}

// UpdateShow func update a Show.
//
//	@Description	update Show
//	@Summary		update a Show
//	@Tags			Show
//	@Accept			json
//	@Produce		json
//	@Param			id					path		string			true	"Show ID"
//	@Param			updateShow			body		schema.UpsertShow	true	"Update a Show"
//	@Success		200					{object}	schema.Show
//	@Failure		400,401,403,404,500	{object}	schema.ErrorResponse	"Error"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/Shows/{id} [put]
func UpdateShow(c *fiber.Ctx) error {
	ID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	show := model.Show{ID: ID}
	err = db.First(&show).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "Show not found",
		})
	}

	updateShow := &schema.UpsertShow{}
	if err := c.BodyParser(updateShow); err != nil {
		// Return 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	// Create a new validator for a Show model.
	validate := validator.NewValidator()
	if err := validate.Struct(updateShow); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":    "invalid input found",
			"errors": validator.ValidatorErrors(err),
		})
	}

	if err := convert.Update(&show, &updateShow); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	if err := db.Save(&show).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.JSON(convert.To[schema.Show](show))
}

// DeleteShow func delete a Show.
//
//	@Description	delete Show
//	@Summary		delete a Show
//	@Tags			Show
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string					true	"Show ID"
//	@Success		200				{object}	interface{}				"Ok"
//	@Failure		401,403,404,500	{object}	schema.ErrorResponse	"Error"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/Shows/{id} [delete]
func DeleteShow(c *fiber.Ctx) error {
	ID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	show := model.Show{ID: ID}
	err = db.First(&show).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "user not found",
		})
	}

	show = model.Show{ID: ID}
	result := db.Delete(&show)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": result.Error,
		})
	}

	return c.JSON(fiber.Map{})
}

// GetShowsByMovieID func gets all shows with the same movieID.
//
//	@Description	get shows by movieID.
//	@Summary		get shows by movieID
//	@Tags			Show
//	@Accept			json
//	@Produce		json
//	@Param			movieID		path		string	true	"Movie ID"
//	@Success		200		{object}	[]schema.Show
//	@Failure		400,404	{object}	schema.ErrorResponse	"Error"
//	@Router			/api/v1/Shows/{movieID} [get]
func GetShowsByMovieID(c *fiber.Ctx) error {
	movieID := c.Params("movieID")
	shows := []model.Show{}
	err := db.Where("movie_id = ?", movieID).Find(&shows).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "Shows not found",
		})
	}
	return c.JSON(convert.To[schema.Show](shows))
}
