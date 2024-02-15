package controller

import (
	"github.com/CSCI-X050-A7/backend/app/model"
	"github.com/CSCI-X050-A7/backend/app/schema"
	"github.com/CSCI-X050-A7/backend/pkg/convert"
	"github.com/CSCI-X050-A7/backend/pkg/validator"
	"github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// GetMovies func gets all movies.
// @Description Get all movies.
// @Summary get all movies
// @Tags Movie
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Success 200 {object} schema.MovieListResponse
// @Failure 400 {object} schema.ErrorResponse "Error"
// @Router /api/v1/movies [get]
func GetMovies(c *fiber.Ctx) error {
	pagination := GetPagination(c)
	objs, count, err := ListObjs[schema.Movie](
		db.Model(model.Movie{}), pagination,
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

// GetMovie func gets a movie.
// @Description a movie.
// @Summary get a movie
// @Tags Movie
// @Accept json
// @Produce json
// @Param id path string true "Movie ID"
// @Success 200 {object} schema.Movie
// @Failure 400,404 {object} schema.ErrorResponse "Error"
// @Router /api/v1/movies/{id} [get]
func GetMovie(c *fiber.Ctx) error {
	ID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	movie := model.Movie{ID: ID}
	err = db.First(&movie).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "movie not found",
		})
	}
	return c.JSON(convert.To[schema.Movie](movie))
}

// CreateMovie func for creates a new movie.
// @Description Create a new movie.
// @Summary create a new movie
// @Tags Movie
// @Accept json
// @Produce json
// @Param createmovie body schema.CreateMovie true "Create new movie"
// @Failure 400,401,500 {object} schema.ErrorResponse "Error"
// @Success 200 {object} schema.Movie "Ok"
// @Security ApiKeyAuth
// @Router /api/v1/movies [post]
func CreateMovie(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID, ok := claims["user_id"]
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "can't extract user info from request",
		})
	}

	// Create new Movie struct
	createMovie := &schema.CreateMovie{}
	if err := c.BodyParser(createMovie); err != nil {
		// Return 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	logrus.Infof("create movie: %+v", createMovie)

	var err error
	createMovie.UserID, err = uuid.Parse(userID.(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	createMovie.Status = 1 // Active

	// Create a new validator for a Movie model.
	validate := validator.NewValidator()
	if err := validate.Struct(createMovie); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":    "invalid input found",
			"errors": validator.ValidatorErrors(err),
		})
	}

	newMovie := model.Movie{}
	if err := convert.Update(&newMovie, &createMovie); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	if err := db.Create(&newMovie).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.JSON(convert.To[schema.Movie](newMovie))
}

// UpdateMovie func update a movie.
// @Description update movie
// @Summary update a movie
// @Tags Movie
// @Accept json
// @Produce json
// @Param id path string true "Movie ID"
// @Param updatemovie body schema.Movie true "Update a movie"
// @Success 200 {object} schema.Movie
// @Failure 400,401,403,404,500 {object} schema.ErrorResponse "Error"
// @Security ApiKeyAuth
// @Router /api/v1/movies/{id} [put]
func UpdateMovie(c *fiber.Ctx) error {
	ID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	movie := model.Movie{ID: ID}
	err = db.First(&movie).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "movie not found",
		})
	}

	updateMovie := &schema.UpdateMovie{}
	if err := c.BodyParser(updateMovie); err != nil {
		// Return 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	// Create a new validator for a Movie model.
	validate := validator.NewValidator()
	if err := validate.Struct(updateMovie); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":    "invalid input found",
			"errors": validator.ValidatorErrors(err),
		})
	}

	if err := convert.Update(&movie, &updateMovie); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	if err := db.Save(&movie).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.JSON(convert.To[schema.Movie](movie))
}

// DeleteMovie func delete a movie.
// @Description delete movie
// @Summary delete a movie
// @Tags Movie
// @Accept json
// @Produce json
// @Param id path string true "Movie ID"
// @Success 200 {object} interface{} "Ok"
// @Failure 401,403,404,500 {object} schema.ErrorResponse "Error"
// @Security ApiKeyAuth
// @Router /api/v1/movies/{id} [delete]
func DeleteMovie(c *fiber.Ctx) error {
	ID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	movie := model.Movie{ID: ID}
	err = db.First(&movie).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "user not found",
		})
	}

	movie = model.Movie{ID: ID}
	result := db.Delete(&movie)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.JSON(fiber.Map{})
}