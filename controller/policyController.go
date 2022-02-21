package controller

import (
	"crypto/rand"
	"encoding/hex"

	email "github.com/floydjones1/auth-server/Email"
	"github.com/floydjones1/auth-server/database"
	"github.com/floydjones1/auth-server/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

//This function is to generate passord reset token for users
func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func PolicyDetails(c *fiber.Ctx) error {
	user:=c.Locals("user").(*jwt.Token)
	claims:=user.Claims.(jwt.MapClaims)
	id:=claims["user_id"].(float64)


	//Check if email already exist in database
	var userData model.User
	database.DB.Where("id=?", id).First(&userData)
	val, _ := randomHex(4)
	
	identity := &model.VerifyCode{
		UserID: float64(userData.Id),
		IdentityCode: val,

	}

	// database.DB.Create(identity)
	email.SendEmailToken(userData.Email,val)
	
	return c.JSON(fiber.Map{
		"message":"Kindly follow the process on the mobile for the capturing.",
		"id":identity,
		 
	})
}