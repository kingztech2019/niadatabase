package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kingztech2019/nia_backend/database"
	"github.com/kingztech2019/nia_backend/model"
)


func GetStatus(c *fiber.Ctx) error {
	verifyCode:= c.Query("id")
	 var uploadstatus model.UploadStatus
	database.DB.Model(&uploadstatus).Where("identity_code", verifyCode).Find(&uploadstatus)
	c.Status(200)
	return c.JSON(fiber.Map{
		"status":uploadstatus,
	})


	
}