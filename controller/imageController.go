package controller

import (
	"encoding/base64"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kingztech2019/nia_backend/database"
	"github.com/kingztech2019/nia_backend/model"
)


type ImageRequest struct {
	FirstImage string  
  
	SecondImage string  
	ThirdImage  string  
	FourthImage string  
}
var fileName string
var firstImage,secondImage,thirdImage,fourthImage string

//This function convert base64 to image
func imageReciever(image string) string{
	 
	dec, err := base64.StdEncoding.DecodeString(image)
	dirs := []string{"", "uploads"}
	for _, dir := range dirs {
	   f, err := os.CreateTemp(dir, "image-*.png")
	   if err != nil {
		  panic(err)
	   }
	   defer f.Close()
	    

	   if _, err := f.Write(dec); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}
	fileName=f.Name()
	
	// log.Println("NAME",f.Name())
	// log.Println("http://localhost:4000/public/"+f.Name())

	// // go to begginng of file
	// f.Seek(0, 0)

	// // output file contents
	// io.Copy(os.Stdout, f)
	}
	 
	// f, err := os.Create(fmt.Sprintf("./utils/%s-fileme.png",randSeq(4)))
	if err != nil {
		panic(err)
	}
	// defer f.Close()

	return fileName
}

func UploadImage(c *fiber.Ctx) error {
	verifyCode:= c.Query("id")
	 
	req := new(ImageRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	// ||req.ThirdImage==""||req.FourthImage=="" 

	if  req.FirstImage == "" || req.SecondImage == ""{
		return fiber.NewError(fiber.StatusBadRequest, "invalid images credentials")
	}
	
	
//Get the user verification code
var identitycode model.VerifyCode
database.DB.Where("identity_code=?", strings.TrimSpace(verifyCode)).Find(&identitycode)
if identitycode.UserID ==0{
	c.Status(404)
	 return c.JSON(fiber.Map{
	  "message": "user doesn't exist",
	})
	
	
  }
  if req.FirstImage !="" {
	firstImage = imageReciever(req.FirstImage)

	
}
if req.SecondImage !="" {
	secondImage = imageReciever(req.SecondImage)

	
}
if req.ThirdImage !="" {
	thirdImage = imageReciever(req.ThirdImage)

	
}
if req.FourthImage !="" {
	fourthImage = imageReciever(req.FourthImage)

	
}
var uploadstatus model.UploadStatus
 
database.DB.Model(&uploadstatus).Where("identity_code=?",  strings.TrimSpace(verifyCode)).Update("upload_status", "active")

  images := &model.ImagesUrl{
	 FirstImage: firstImage,
	 SecondImage: secondImage,
	 ThirdImage: thirdImage,
	 FourthImage: fourthImage,
	 UserID: identitycode.UserID,
}

database.DB.Create(images)
    
return c.JSON(fiber.Map{"message": "Image uploaded Successfully"})
}