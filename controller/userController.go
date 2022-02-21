package controller

import (
	"regexp"
	"strings"

	"github.com/floydjones1/auth-server/database"
	"github.com/floydjones1/auth-server/model"
	"github.com/floydjones1/auth-server/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	//Name     string
	Email    string
	Password string
}

type LoginRequest struct {
	Email    string
	Password string
}

func validateEmail(email string) bool{
	
	Re:= regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
	return Re.MatchString(email)
}

func SignUp(c *fiber.Ctx) error {
	var userData model.User
	req := new(SignupRequest)
		if err := c.BodyParser(req); err != nil {
			return err
		}

		if  req.Email == "" || req.Password == "" {
			return fiber.NewError(fiber.StatusBadRequest, "invalid signup credentials")
		}

		//Check if email already exist in database
	database.DB.Where("email=?", strings.TrimSpace(req.Email)).First(&userData)
	if userData.Id!=0{
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"Email already exist",
		})

	}

	if !validateEmail(strings.TrimSpace(req.Email)){
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"Invalid Email Address",
		})

	}
	//Check if password is less than 6 characters
	if len(req.Password)<=6{
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"Password must be greater than 6 character",
		})
	}


		// save this info in the database
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user := &model.User{
			 
			Email:    req.Email,
			Password: string(hash),
		}
 
		 database.DB.Create(user)
		if err != nil {
			return err
		}
		token, exp, err := utils.CreateJWTToken(*user)
		if err != nil {
			return err
		}
		// create a jwt token

		return c.JSON(fiber.Map{"token": token, "exp": exp, "user": user})
	
}

func Login(c *fiber.Ctx) error {
	req := new(LoginRequest)
		if err := c.BodyParser(req); err != nil {
			return err
		}

		if req.Email == "" || req.Password == "" {
			return fiber.NewError(fiber.StatusBadRequest, "invalid login credentials")
		}

		user := new(model.User)
		database.DB.Where("email=?", req.Email).First(&user) 
	if user.Id ==0{
		c.Status(404)
		return c.JSON(fiber.Map{
			"message":"Email Address doesn't exit, kindly create an account",
		})
	}
 		 
		// if !has {
		// 	return fiber.NewError(fiber.StatusBadRequest, "invalid login credentials")
		// }

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			c.Status(400)
		return c.JSON(fiber.Map{
			"message":"incorrect password",
		})
		}

		token, exp, err := utils.CreateJWTToken(*user)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{"token": token, "exp": exp, "user": user})
	
}

// func Home(c *fiber.Ctx) error  {
// 	   return c.JSON(fiber.Map{"success": true, "path": "public"})
	
	
// }

// func Private(c *fiber.Ctx) error { 
// 	user:=c.Locals("user").(*jwt.Token)
// 	claims:=user.Claims.(jwt.MapClaims)
// 	id:=claims["user_id"].(float64)


// 	return c.JSON(fiber.Map{"success": true, "path": "private","id":id })
	
	
// }
