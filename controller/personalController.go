package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kingztech2019/nia_backend/database"
	"github.com/kingztech2019/nia_backend/model"
)

type VinCode struct{
	Vin string `json:vin`
}
type Post struct {
	Make string `json:"make"`
	Model     string    `json:"model"`
	Year  string `json:"year"`
	engine   string `json:"engine"`
   }

   type Data struct {
	 Value int `json:"value"`
   }

// func TestTrue(value string,c fiber.Ctx, model string,year string  ) error {
// 	log.Println("THIS IS VALUE",value)
// 	postData, _ := json.Marshal(map[string]string{
// 		"make":value,
// 		"year":  year ,
// 		"model": model,
// 	 })
	  
// 	 responseData := bytes.NewBuffer(postData) 
	
//   dataresp, err := http.Post("https://truevalue.octamile.com/analysis2","application/json",
//   responseData)
//   log.Println("LLOOK",responseData)
//    if err != nil {  
	   
// 	 log.Printf("Request Failed: %s", err)
// 	 return nil
//    }

//    defer dataresp.Body.Close()
//    dataBody, err := ioutil.ReadAll(dataresp.Body)
//    // Log the request body 
//    databodyString := string(dataBody)
//    log.Print(databodyString)
//    // Unmarshal result
//     data:=Data{}
//    err = json.Unmarshal(dataBody, &data)
//    if err != nil {
// 	  log.Printf("Reading body failed: %s", err)
// 	  return nil
//    }
//    log.Println("BEED",data.Value)
// //    return c.JSON(fiber.Map{
// // 	   "message":data.Value,
// //    })
   

// return nil
	
// }
func CheckVin(c *fiber.Ctx) error {
	req := new(VinCode)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	if  req.Vin == "" {
		return fiber.NewError(fiber.StatusBadRequest, "invalid signup credentials")
	}
	postBody, _ := json.Marshal(map[string]string{
		"vin":req.Vin,
		"email":  "tester",
		"password": "tester",
	 })
	 responseBody := bytes.NewBuffer(postBody)
  resp, err := http.Post("https://autorescue.ng/dealer/v2/vinprev","application/json",
  responseBody)
   if err != nil {   
	 log.Printf("Request Failed: %s", err)
	 return nil
   }
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// Log the request body 
	bodyString := string(body)
	log.Print(bodyString)
	// Unmarshal resultresponseData
	post := Post{}
	err = json.Unmarshal(body, &post)
	if err != nil {
	   log.Printf("Reading body failed: %s", err)
	   return nil
	}
	postData, _ := json.Marshal(map[string]string{
		"make":post.Make,
		"year":  post.Year ,
		"model": post.Model,
	 })
	  
	 responseData := bytes.NewBuffer(postData) 
	
  dataresp, err := http.Post("https://truevalue.octamile.com/analysis2","application/json",
  responseData)
  log.Println("LLOOK",responseData)
   if err != nil {  
	   
	 log.Printf("Request Failed: %s", err)
	 return nil
   }

   defer dataresp.Body.Close()
   dataBody, err := ioutil.ReadAll(dataresp.Body)
   // Log the request body 
   databodyString := string(dataBody)
   log.Print(databodyString)
   // Unmarshal result
    data:=Data{}
   err = json.Unmarshal(dataBody, &data)
   if err != nil {
	  log.Printf("Reading body failed: %s", err)
	  return nil
   }
 
	return c.JSON(fiber.Map{
		"data":fiber.Map{
			"firstcheck":bodyString,
			"finalcheck":data.Value,
		},
			
		
 	})
	//return nil 
	 
	
}
 

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