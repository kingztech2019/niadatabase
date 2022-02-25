package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kingztech2019/nia_backend/database"
	"github.com/kingztech2019/nia_backend/model"
)


type VechicleRequest struct {
	ClassOfInsurance    string
	Type string
	VechicleUse string
}
func VechicleInsurance(c *fiber.Ctx) error {
	req := new(VechicleRequest)
		if err := c.BodyParser(req); err != nil {
			return err
		}

		if req.ClassOfInsurance == "" || req.Type == ""||req.VechicleUse=="" {
			return fiber.NewError(fiber.StatusBadRequest, "invalid vechicle insurance credentials")
		}
		user:=c.Locals("user").(*jwt.Token)
		claims:=user.Claims.(jwt.MapClaims)
		id:=claims["user_id"].(float64)
			log.Println(id)
		vechicleDetail := &model.VechicleInsurance{
			  UserID: id,
			  ClassOfInsurance: req.ClassOfInsurance,
			  Type: req.Type,
			  VechicleUse: req.VechicleUse,
		}
		database.DB.Create(vechicleDetail)
		return c.JSON(fiber.Map{"message":"Your Order is Successful"})

	
}