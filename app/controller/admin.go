package controller

import (
	"github.com/CSCI-X050-A7/backend/app/model"
	"github.com/CSCI-X050-A7/backend/app/schema"
	"github.com/CSCI-X050-A7/backend/pkg/convert"
	"github.com/CSCI-X050-A7/backend/pkg/validator"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

// GetUsers func get all users.
//
//	@Description	Get all users.
//	@Summary		get all users
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			offset		query		integer	false	"offset"
//	@Param			limit		query		integer	false	"limit"
//	@Success		200			{object}	schema.UserDetailListResponse
//	@Failure		400,401,403	{object}	schema.ErrorResponse	"Error"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/admin/users [get]
func AdminGetUsers(c *fiber.Ctx) error {
	pagination := GetPagination(c)
	objs, count, err := ListObjs[schema.UserDetail](
		db.Model(model.User{}), pagination,
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

// GetUser func get a user.
//
//	@Description	a user.
//	@Summary		get a user
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"User ID"
//	@Success		200				{object}	schema.UserDetail
//	@Failure		400,401,403,404	{object}	schema.ErrorResponse	"Error"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/admin/users/{id} [get]
func AdminGetUser(c *fiber.Ctx) error {
	ID, err := uuid.Parse(c.Params("id"))
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
	return c.JSON(convert.To[schema.UserDetail](user))
}

// CreateUser func for creates a new user.
//
//	@Description	Create a new user.
//	@Summary		create a new user
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			createuser		body		schema.CreateUser		true	"Create new user"
//	@Failure		400,401,409,500	{object}	schema.ErrorResponse	"Error"
//	@Success		200				{object}	schema.User				"Ok"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/admin/users [post]
func AdminCreateUser(c *fiber.Ctx) error {
	// Create new User struct
	createUser := &schema.CreateUser{}

	if err := c.BodyParser(createUser); err != nil {
		// Return 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	// Create a new validator for a User model.
	validate := validator.NewValidator()
	if err := validate.Struct(createUser); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":    "invalid input found",
			"errors": validator.ValidatorErrors(err),
		})
	}

	// check user already exists
	result := db.Where(model.User{Email: createUser.Email}).
		Or(model.User{UserName: createUser.UserName}).Find(&model.User{})
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": result.Error.Error(),
		})
	}
	if result.RowsAffected != 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"msg": "user with this username or email already exists",
		})
	}

	createUser.Password, _ = GeneratePasswordHash([]byte(createUser.Password))
	newUser := model.User{}
	if err := convert.Update(&newUser, &createUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	if err := db.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.JSON(convert.To[schema.User](newUser))
}

// UpdateUser func update a user.
//
//	@Description	first_name, last_name, is_active, is_admin only
//	@Summary		update a user
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			id					path		string				true	"User ID"
//	@Param			updateuser			body		schema.AdminUpdateUser	true	"Update a user"
//	@Success		200					{object}	schema.User
//	@Failure		400,401,403,404,500	{object}	schema.ErrorResponse	"Error"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/admin/users/{id} [put]
func AdminUpdateUser(c *fiber.Ctx) error {
	ID, err := uuid.Parse(c.Params("id"))
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

	updateUser := &schema.AdminUpdateUser{}
	if err := c.BodyParser(updateUser); err != nil {
		// Return 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	// Create a new validator for a User model.
	validate := validator.NewValidator()
	if err := validate.Struct(updateUser); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":    "invalid input found",
			"errors": validator.ValidatorErrors(err),
		})
	}

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

	return c.JSON(convert.To[schema.User](user))
}

// DeleteUser func delete a user.
//
//	@Description	delete user
//	@Summary		delete a user
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string					true	"User ID"
//	@Success		200				{object}	interface{}				"Ok"
//	@Failure		401,403,404,500	{object}	schema.ErrorResponse	"Error"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/admin/users/{id} [delete]
func AdminDeleteUser(c *fiber.Ctx) error {
	ID, err := uuid.Parse(c.Params("id"))
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

	user = model.User{ID: ID}
	result := db.Delete(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": result.Error,
		})
	}

	return c.JSON(fiber.Map{})
}
