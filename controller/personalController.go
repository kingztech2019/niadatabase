package controller

import (
	"fmt"

	"github.com/floydjones1/auth-server/database"
	"github.com/floydjones1/auth-server/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)


func PersonalDetails(c *fiber.Ctx) error {
	var data map[string]string
	
	if err:=c.BodyParser(&data); err != nil {
	  fmt.Println("Unable to parse body")
	}
	user:=c.Locals("user").(*jwt.Token)
	claims:=user.Claims.(jwt.MapClaims)
	id:=claims["user_id"].(float64)

	personalData:=model.PersonalDetails{
		UserID: id,
		FirstName: data["first_name"],
		LastName: data["last_name"],
		Id: data["id"],
		MeansOfId: data["means_of_id"],
		MiddleName: data["middle_name"],
		PhoneNumber: data["phone_number"],
		Email: data["email"],
		State: data["state"],
		Lga: data["lga"],
		Address: data["address"],
		Title: data["title"],

	}

	 database.DB.Create(&personalData)
	 return c.JSON(fiber.Map{
		 "Message":"Details save successfully",
	 })
	
}