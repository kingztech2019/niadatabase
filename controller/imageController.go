package controller

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kingztech2019/nia_backend/database"
	"github.com/kingztech2019/nia_backend/model"
)

type ImageRequest struct {
	FirstImage string

	SecondImage string
	ThirdImage  string
	FourthImage string
}

type Analysis struct {
	LabelModelVersion string `json:"labelmodelversion"`
	Labels            []struct {
		Confidence float64 `json:"confidence"`
		Name       string  `json:"name"`
		Parents    []struct {
			Name string `json:"name"`
		}
	}
}

var fileName string
var filtered = []string{}
var data []string
var firstImage, secondImage, thirdImage, fourthImage string
var firstAnalysis, secondAnalysis, thirdAnalysis, fourthAnalysis []byte

//THIS FUNCTION CONVERT THE IMAGES TO A CORRECT JSON FORMAT
func ImageValidator(scannedimage []byte) interface{} {
	var dat Analysis
	if err := json.Unmarshal(scannedimage, &dat); err != nil {
		panic(err)
	}
	result := dat
	for _, value := range result.Labels {
		for _, val := range value.Parents {

			if val.Name == "Car" && value.Confidence >= float64(60) || val.Name == "Vehicle" && value.Confidence >= float64(60) {
				filter := append(filtered, val.Name)
				//fmt.Println(filter)
				return filter

			}

		}
	}

	return nil
}

//THIS FUNCTION USE AMAZON IMAGE RECOGNITIONTO DETECT THE IMAGE
func imageScanner(image string) []byte {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}
	sess := session.New(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AccessKeyID"),
			os.Getenv("SecretAccessKey"),
			""),
	})

	svc := rekognition.New(sess)

	decodedImage, err := base64.StdEncoding.DecodeString(image)

	if err != nil {

		return []byte(err.Error())
	}

	// Send request to Rekognition.
	input := &rekognition.DetectLabelsInput{
		Image: &rekognition.Image{
			Bytes: decodedImage,
		},
	}

	result, err := svc.DetectLabels(input)
	if err != nil {
		return []byte(err.Error())
	}
	output, err := json.Marshal(result)
	if err != nil {

		return []byte(err.Error())
	}
	//log.Println(output)
	return output

}

//This function convert base64 to image
func imageReciever(image string) string {

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
		fileName = f.Name()

	}

	if err != nil {
		panic(err)
	}
	// defer f.Close()

	return fileName
}

//THIS FUNCTION RETURN THE FINAL RESULT
func UploadImage(c *fiber.Ctx) error {
	verifyCode := c.Query("id")

	req := new(ImageRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	// ||req.ThirdImage==""||req.FourthImage==""

	if req.FirstImage == "" || req.SecondImage == "" {
		return fiber.NewError(fiber.StatusBadRequest, "invalid images payload")
	}

	//Get the user verification code
	var identitycode model.VerifyCode
	database.DB.Where("identity_code=?", strings.TrimSpace(verifyCode)).Find(&identitycode)
	if identitycode.UserID == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "user doesn't exist",
		})

	}
	if req.FirstImage != "" {
		firstAnalysis = imageScanner(req.FirstImage)
		if ImageValidator(firstAnalysis) == nil {
			c.Status(400)
			return c.JSON(fiber.Map{
				"message": "Unable to identify the first image uploaded please try again",
			})
		}

	}
	if req.SecondImage != "" {

		secondAnalysis = imageScanner(req.SecondImage)
		if ImageValidator(secondAnalysis) == nil {
			c.Status(400)
			return c.JSON(fiber.Map{
				"message": "Unable to identify the second image uploaded please try again",
			})
		}

	}
	if req.ThirdImage != "" {
		thirdAnalysis = imageScanner(req.ThirdImage)
		if ImageValidator(thirdAnalysis) == nil {
			c.Status(400)
			return c.JSON(fiber.Map{
				"message": "Unable to identify the third image uploaded please try again",
			})
		}

	}
	if req.FourthImage != "" {
		fourthAnalysis = imageScanner(req.FourthImage)
		if ImageValidator(fourthAnalysis) == nil {
			c.Status(400)
			return c.JSON(fiber.Map{
				"message": "Unable to identify the fourth image uploaded please try again",
			})
		}

	}

	//THIS METHOD CREATE AND SAVE THE IMAGES
	firstImage = imageReciever(req.FirstImage)
	secondImage = imageReciever(req.SecondImage)
	thirdImage = imageReciever(req.ThirdImage)
	fourthImage = imageReciever(req.FourthImage)
	var uploadstatus model.UploadStatus
	var user model.User
	var personalData model.PersonalDetails
	database.DB.Where("user_id=?", identitycode.UserID).First(&personalData)

	database.DB.Model(&uploadstatus).Where("identity_code=?", strings.TrimSpace(verifyCode)).Update("upload_status", "active")
	database.DB.Where("id=?", identitycode.UserID).First(&user)
	// if uploadstatus.UploadStatus == "active" {
	// 	email.SendCertificateMail(user.Email, personalData.FirstName)

	// }

	images := &model.ImagesUrl{
		FirstImage:  firstImage,
		SecondImage: secondImage,
		ThirdImage:  thirdImage,
		FourthImage: fourthImage,
		UserID:      identitycode.UserID,
	}

	database.DB.Create(images)

	return c.JSON(fiber.Map{
		"message": "Image uploaded successfully",
	})
}
