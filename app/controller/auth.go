package controller

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/CSCI-X050-A7/backend/app/model"
	"github.com/CSCI-X050-A7/backend/app/schema"
	"github.com/CSCI-X050-A7/backend/pkg/config"
	"github.com/CSCI-X050-A7/backend/pkg/convert"
	"github.com/CSCI-X050-A7/backend/pkg/email"
	"github.com/CSCI-X050-A7/backend/pkg/validator"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Login method for creating a new access token.
//
//	@Description	Set new access token to cookies and redirect. Demo username: demo, password: 123456
//	@Summary		login
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			login			body		schema.Auth				true	"Request for token"
//	@Param			redirect_url	query		string					false	"Redirect url after login"
//	@Failure		400,404,401,500	{object}	schema.ErrorResponse	"Error"
//	@Success		200				{object}	schema.TokenResponse	"Ok"
//	@Router			/api/v1/auth/login [post]
func Login(c *fiber.Ctx) error {
	login := &schema.Auth{}
	if err := c.BodyParser(login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	redirect_url := c.Query("redirect_url", "/")
	user := model.User{}
	err := db.Where(&model.User{UserName: login.Username}).First(&user).Error
	logrus.Infof("user: %v", user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "username not found",
		})
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password),
		[]byte(login.Password)) != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"msg": "password is wrong",
		})
	}

	if !user.IsActive {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"msg": "user not active anymore.",
		})
	}

	// Generate a new Access token.
	token, err := GenerateNewAccessToken(user.ID, user.IsAdmin)
	if err != nil {
		// Return 500 and token generation error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:  "access_token",
		Value: token,
		Expires: time.Now().
			Add(time.Duration(config.Conf.JWTExpireSeconds) * time.Second),
		HTTPOnly: true,
		SameSite: "lax",
	})
	return c.JSON(fiber.Map{
		"msg": fmt.Sprintf("Token will be expired within %d seconds",
			config.Conf.JWTExpireSeconds),
		"access_token": token,
		"redirect_url": redirect_url,
	})
}

// Register method for creating a new unactivated user.
//
//	@Description	Register for a new user.
//	@Summary		register
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			register		body		schema.RegisterUser		true	"register user"
//	@Failure		400,404,401,500	{object}	schema.ErrorResponse	"Error"
//	@Success		200				{object}	schema.User	"Ok"
//	@Router			/api/v1/auth/register [post]
func Register(c *fiber.Ctx) error {
	register := &schema.RegisterUser{}
	if err := c.BodyParser(register); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	validate := validator.NewValidator()
	if err := validate.Struct(register); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":    "invalid input found",
			"errors": validator.ValidatorErrors(err),
		})
	}
	// check user already exists
	result := db.Where(model.User{Email: register.Email}).
		Or(model.User{UserName: register.UserName}).Find(&model.User{})
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
	register.Password, _ = GeneratePasswordHash([]byte(register.Password))
	newUser := model.User{}
	if err := convert.Update(&newUser, &register); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	val, err := rand.Int(rand.Reader, big.NewInt(100000))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	newUser.ActivationCode = fmt.Sprintf("%06d", val)
	if err := db.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	go func() {
		err := email.Send(
			newUser.Email,
			"Cinema E-Booking System Register Confirmation",
			fmt.Sprintf("Link to activate: %s/activate?id=%s&code=%s",
				config.Conf.Url, newUser.ID, newUser.ActivationCode),
		)
		if err != nil {
			logrus.Errorf("email send error: %v", err)
		}
	}()
	return c.JSON(convert.To[schema.User](newUser))
}

// Activation method for activating a new unactivated user.
//
//	@Description	Activate a new user.
//	@Summary		activate
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			id		query		string	false	"id"
//	@Param			code	query		string	false	"code"
//	@Failure		400,404,401,500	{object}	schema.ErrorResponse	"Error"
//	@Success		200				{object}	schema.User	"Ok"
//	@Router			/api/v1/auth/activate [post]
func Activate(c *fiber.Ctx) error {
	code := c.Query("code")
	ID, err := uuid.Parse(c.Query("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	user := model.User{ID: ID}
	err = db.First(&user).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "user not found",
		})
	}
	if user.IsActive {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "user already activated",
		})
	}
	if user.ActivationCode != code {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "activation code is wrong",
		})
	}
	user.IsActive = true
	if err := db.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	return c.JSON(convert.To[schema.User](user))
}

// Logout method.
//
//	@Description	Clean cookies
//	@Summary		Logout
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Failure		400,404,401,500	{object}	schema.ErrorResponse	"Error"
//	@Success		200				{object}	interface{}				"Ok"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/auth/logout [post]
func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name: "access_token",
		// Set expiry date to the past
		Expires:  time.Now().Add(-(time.Hour * 2)),
		HTTPOnly: true,
		SameSite: "lax",
	})
	return nil
}

// Get current JWT method for debugging.
//
//	@Description	Get current JWT.
//	@Summary		JWT
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Failure		400,404,401,500	{object}	schema.ErrorResponse	"Error"
//	@Success		200				{object}	schema.JWT				"Ok"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/auth/jwt [post]
func JWT(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return c.JSON(claims)
}

func GenerateNewAccessToken(userID uuid.UUID, isAdmin bool) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["admin"] = isAdmin
	claims["exp"] = time.Now().Add(time.Duration(config.Conf.JWTExpireSeconds) *
		time.Second).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.Conf.JWTSecret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func GeneratePasswordHash(password []byte) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(password,
		bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func IsValidPassword(hash, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}

// ForgotPassword method for initiating the password reset process.
//
//	@Description	Initiate the password reset process.
//	@Summary		forgot password
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			email	query		string	false	"Email"
//	@Failure		400,404,401,500	{object}	schema.ErrorResponse	"Error"
//	@Success		200				{object}	interface{}				"Ok"
//	@Router			/api/v1/auth/forgotpassword [post]
func ForgotPassword(c *fiber.Ctx) error {
	email := c.Query("email")
	user := model.User{}
	err := db.Where(&model.User{Email: email}).First(&user).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "user not found",
		})
	}

	// Generate a random password reset code
	passwordCode, err := generateRandomString(10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "failed to generate password reset code",
		})
	}

	// Update the user's PasswordCode field with the generated code
	user.PasswordCode = passwordCode
	if err := db.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "failed to save password reset code",
		})
	}

	// Send the password reset link to the user's emai
	go func() {
		//err := email.Send(
		//user.Email,
		//"Cinema E-Booking System Password Reset",
		//fmt.Sprintf("Link to reset password: %s/reset-password?id=%s&code=%s",
		//	config.Conf.Url, user.ID, user.PasswordCode),
		//)
		if err != nil {
			logrus.Errorf("email send error: %v", err)
		}
	}()

	return c.JSON(fiber.Map{
		"msg": "password reset link sent to your email",
	})
}

// ResetPassword method for resetting the user's password.
//
//	@Description	Reset the user's password.
//	@Summary		reset password
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			id		query		string	false	"id"
//	@Param			code	query		string	false	"code"
//	@Param			password	body		string	true	"New password"
//	@Failure		400,404,401,500	{object}	schema.ErrorResponse	"Error"
//	@Success		200				{object}	interface{}				"Ok"
//	@Router			/api/v1/auth/resetpassword [post]
func ResetPassword(c *fiber.Ctx) error {
	code := c.Query("code")
	ID, err := uuid.Parse(c.Query("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	user := model.User{ID: ID}
	err = db.First(&user).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "user not found",
		})
	}
	if user.PasswordCode != code {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "invalid password reset code",
		})
	}

	// Update the user's password with the new password
	newPassword := c.FormValue("password")
	hashedPassword, err := GeneratePasswordHash([]byte(newPassword))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "failed to generate password hash",
		})
	}
	user.Password = hashedPassword
	user.PasswordCode = "" // Clear the password reset code
	if err := db.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "failed to save new password",
		})
	}

	return c.JSON(fiber.Map{
		"msg": "password reset successful",
	})
}

func generateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	for i := 0; i < length; i++ {
		b[i] = charset[int(b[i])%len(charset)]
	}
	return string(b), nil
}
