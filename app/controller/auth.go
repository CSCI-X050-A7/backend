package controller

import (
	"fmt"
	"time"

	"github.com/CSCI-X050-A7/backend/app/model"
	"github.com/CSCI-X050-A7/backend/app/schema"
	"github.com/CSCI-X050-A7/backend/pkg/config"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Login method for creating a new access token.
// @Description Set new access token to cookies and redirect. Demo username: demo, password: 123456
// @Summary login
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body schema.Auth true "Request for token"
// @Param redirect_url query string false "Redirect url after login"
// @Failure 400,404,401,500 {object} schema.ErrorResponse "Error"
// @Success 200 {object} schema.TokenResponse "Ok"
// @Router /api/v1/auth/login [post]
func Login(c *fiber.Ctx) error {
	login := &schema.Auth{}
	if err := c.BodyParser(login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	redirect_url := c.Params("redirect_url", "/")
	user := model.User{UserName: login.Username}
	err := db.First(&user).Error
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

// Logout method.
// @Description Clean cookies
// @Summary Logout
// @Tags Auth
// @Accept json
// @Produce json
// @Failure 400,404,401,500 {object} schema.ErrorResponse "Error"
// @Success 200 {object} interface{} "Ok"
// @Security ApiKeyAuth
// @Router /api/v1/auth/logout [post]
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
// @Description Get current JWT.
// @Summary JWT
// @Tags Auth
// @Accept json
// @Produce json
// @Failure 400,404,401,500 {object} schema.ErrorResponse "Error"
// @Success 200 {object} interface{} "Ok"
// @Security ApiKeyAuth
// @Router /api/v1/auth/jwt [post]
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
