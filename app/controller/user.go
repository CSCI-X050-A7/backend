package controller

import (
	"github.com/CSCI-X050-A7/backend/app/model"
	"github.com/CSCI-X050-A7/backend/app/schema"
	"github.com/CSCI-X050-A7/backend/pkg/config"
	"github.com/CSCI-X050-A7/backend/pkg/convert"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

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
	userDetail := convert.To[schema.UserDetail](user)
	key := []byte(config.Conf.JWTSecret)
	userDetail.CardType, _ = AESDecrypt(key, userDetail.CardType)
	userDetail.CardNumber, _ = AESDecrypt(key, userDetail.CardNumber)
	userDetail.CardExpiration, _ = AESDecrypt(key, userDetail.CardExpiration)

	/* For easy debugging:
	var str1, str2, str3 string
	var err1, err2, err3 error
	str1 = "American Express"
	str2 = "thisismycardnumber"
	str3 = "exp"
	str1, err4 = AESEncrypt(key, str1)
	str2, err5 = AESEncrypt(key, str2)
	str3, err6 = AESEncrypt(key, str3)
	fmt.Println("\n\n\nAfter encrypt")
	fmt.Printf("%v, %v, %v\n", err1, err2, err3)
	fmt.Printf("%s, %s, %s", str1, str2, str3)
	fmt.Println("\n\n\n")
	str1, err4 = AESDecrypt(key, str1)
	str2, err5 = AESDecrypt(key, str2)
	str3, err6 = AESDecrypt(key, str3)

	fmt.Println("\n\n\nAfter decrypt")
	fmt.Printf("%v, %v, %v", err1, err2, err3)
	fmt.Printf("%s, %s, %s", str1, str2, str3)
	fmt.Println("\n\n\n")
	*/

	return c.JSON(userDetail)
}
