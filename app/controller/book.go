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

// GetBooks func gets all books.
// @Description Get all books.
// @Summary get all books
// @Tags Book
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Success 200 {array} schema.Book
// @Failure 400 {object} schema.ErrorResponse "Error"
// @Router /api/v1/books [get]
func GetBooks(c *fiber.Ctx) error {
	pagination := GetPagination(c)
	objs, count, err := ListObjs[schema.Book](
		db.Model(model.Book{}), pagination,
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

// GetBook func gets a book.
// @Description a book.
// @Summary get a book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} schema.Book
// @Failure 400,404 {object} schema.ErrorResponse "Error"
// @Router /api/v1/books/{id} [get]
func GetBook(c *fiber.Ctx) error {
	ID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	book := model.Book{ID: ID}
	err = db.First(&book).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "book not found",
		})
	}
	return c.JSON(convert.To[schema.Book](book))
}

// CreateBook func for creates a new book.
// @Description Create a new book.
// @Summary create a new book
// @Tags Book
// @Accept json
// @Produce json
// @Param createbook body schema.CreateBook true "Create new book"
// @Failure 400,401,500 {object} schema.ErrorResponse "Error"
// @Success 200 {object} schema.Book "Ok"
// @Security ApiKeyAuth
// @Router /api/v1/books [post]
func CreateBook(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID, ok := claims["user_id"]
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "can't extract user info from request",
		})
	}

	// Create new Book struct
	createBook := &schema.CreateBook{}
	if err := c.BodyParser(createBook); err != nil {
		// Return 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	logrus.Infof("create book: %+v", createBook)

	var err error
	createBook.UserID, err = uuid.Parse(userID.(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	createBook.Status = 1 // Active

	// Create a new validator for a Book model.
	validate := validator.NewValidator()
	if err := validate.Struct(createBook); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":    "invalid input found",
			"errors": validator.ValidatorErrors(err),
		})
	}

	newBook := model.Book{}
	if err := convert.Update(&newBook, &createBook); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	if err := db.Create(&newBook).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.JSON(convert.To[schema.Book](newBook))
}

// UpdateBook func update a book.
// @Description update book
// @Summary update a book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param updatebook body schema.Book true "Update a book"
// @Success 200 {object} schema.Book
// @Failure 400,401,403,404,500 {object} schema.ErrorResponse "Error"
// @Security ApiKeyAuth
// @Router /api/v1/books/{id} [put]
func UpdateBook(c *fiber.Ctx) error {
	ID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	book := model.Book{ID: ID}
	err = db.First(&book).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "book not found",
		})
	}

	updateBook := &schema.UpdateBook{}
	if err := c.BodyParser(updateBook); err != nil {
		// Return 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	// Create a new validator for a Book model.
	validate := validator.NewValidator()
	if err := validate.Struct(updateBook); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":    "invalid input found",
			"errors": validator.ValidatorErrors(err),
		})
	}

	if err := convert.Update(&book, &updateBook); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	if err := db.Save(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.JSON(convert.To[schema.Book](book))
}

// DeleteBook func delete a book.
// @Description delete book
// @Summary delete a book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} interface{} "Ok"
// @Failure 401,403,404,500 {object} schema.ErrorResponse "Error"
// @Security ApiKeyAuth
// @Router /api/v1/books/{id} [delete]
func DeleteBook(c *fiber.Ctx) error {
	ID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	book := model.Book{ID: ID}
	err = db.First(&book).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "user not found",
		})
	}

	book = model.Book{ID: ID}
	result := db.Delete(&book)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.JSON(fiber.Map{})
}
