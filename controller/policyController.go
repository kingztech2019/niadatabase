package controller

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	email "github.com/kingztech2019/nia_backend/Email"
	"github.com/kingztech2019/nia_backend/database"
	"github.com/kingztech2019/nia_backend/model"
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
	var data map[string]string
	fmt.Println(data["reg_no"])

	if err:=c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body",err)
		return err
	  }
	user:=c.Locals("user").(*jwt.Token)
	claims:=user.Claims.(jwt.MapClaims)
	id:=claims["user_id"].(float64)


	//Check if email already exist in database
	var userData model.User
	var personalData model.PersonalDetails
	database.DB.Where("id=?", id).First(&userData)
	database.DB.Where("user_id=?",userData.Id).First(&personalData)
	val, _ := randomHex(4)
	
	identity := &model.VerifyCode{
		UserID: float64(userData.Id),
		IdentityCode: val,

	}

	uploadstatus := &model.UploadStatus{
		UploadStatus: "pending",
		UserID: float64(userData.Id),
		IdentityCode: val ,
	}
	database.DB.Create(uploadstatus)
	database.DB.Create(identity)
	email.SendCaptureUrl(userData.Email, personalData.FirstName,val)
	vechicleData:=model.VechicleDetails{
		UserID: id,
		PlateNo: data["plate_no"],
		Vin: data["vin"],
		Engine: data["engine"],
		VechicleColor: data["vechicle_color"],
		Modell: data["model"],
		Value: data["value"],
		Capacity: data["capacity"],
		Make: data["make"],

	}

	policyData:=model.PolicyDetails{
		UserID: id,
		PolicyHolder: data["policy_holder"],
		PhoneNumber: data["phone_number"],
		Email: data["email"],
		Company: data["company"],
		Nin: data["nin"],
		State: data["state"],
		Lga: data["lga"],
		Address: data["address"],

	}
	database.DB.Create(&vechicleData)
	database.DB.Create(&policyData)
	return c.JSON(fiber.Map{
		"message":"Kindly follow the process on the mobile for the capturing.",
		"identity":val,
		 
		 
	})
}