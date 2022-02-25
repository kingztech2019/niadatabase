package controller

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	email "github.com/kingztech2019/nia_backend/Email"
	"github.com/kingztech2019/nia_backend/database"
	"github.com/kingztech2019/nia_backend/model"
	"github.com/kingztech2019/nia_backend/utils"
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

//check if email exist for password reset
func CheckEmailPaswordReset(c *fiber.Ctx) error{
	var user model.User
	var data map[string]string
	if err:=c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body",err)
	  }
	  if !validateEmail(strings.TrimSpace(data["email"])){
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"Invalid Email Address",
		})
	}
		database.DB.Where("email=?", strings.TrimSpace(data["email"])).First(&user)
		if user.Id ==0{
			c.Status(404)
			return c.JSON(fiber.Map{
				"message":"Email Address doesn't exit, kindly use your correct email",
			})
		}

	val, _ := randomHex(4)
     var confirmCode = model.PasswordToken{
       Token: strings.ToUpper(val),
       UserID: uint(user.Id),
       Used: false,
     }
	 email.SendPasswordToken(user.Email,val)
     database.DB.Create(&confirmCode)
	 return c.JSON(fiber.Map{
		 "message":"Password Reset token sent to your email",
		 "token":confirmCode,
	 })

	
}

func ChangePassword(c *fiber.Ctx) error {
	var data map[string]string 
	if err:=c.BodyParser(&data); err != nil {
      fmt.Println("Unable to parse body",err)
    }
	var resetpassword model.PasswordToken
	database.DB.Where("token=?", strings.TrimSpace(data["token"])).Find(&resetpassword)
	if resetpassword.UserID ==0{
		      c.Status(404)
		       return c.JSON(fiber.Map{
		        "message": "Invalid Token",
		 })
	}
	timeCreated:=resetpassword.CreatedAt
    expiredTime:=timeCreated.Add(2 * time.Hour)
    compareDate:=time.Now().After(expiredTime)
    if compareDate{
      //database.DB.Where("user_id = ?", user.Id).Delete(&confirmCode)
      c.Status(404)
      return c.JSON(fiber.Map{
       "message": "Password reset code is expired. ",
     })
	}

	 database.DB.Model(&resetpassword).Where("token = ?", data["token"]).Update("used", true)
	 if len(data["password"])  <= 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
		"message": "Password must be more than 6 character",
	  })
		 
	  }
	  hash, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	  if err != nil {
		  return err
	  }
	  var userData model.User

	//   user:=c.Locals("user").(*jwt.Token)
	// 	claims:=user.Claims.(jwt.MapClaims)
	// 	id:=claims["user_id"].(float64)
	  fmt.Println(resetpassword.UserID)
      database.DB.Model(&userData).Where("id=?",resetpassword.UserID).Update("password", hash)
	  return c.JSON(fiber.Map{
		"message": "Password reset successfully ",
	  })
	
}

// This function update the password 
// func ForgetPassword(c *fiber.Ctx) error {
//     var data map[string]string
    
//     if err:=c.BodyParser(&data); err != nil {
//       fmt.Println("Unable to parse body",err)
//     }
//     var resetpassword model.PasswordToken
//     database.DB.Where("token=?", strings.TrimSpace(data["token"])).Find(&resetpassword)
//     if resetpassword.UserID ==0{
//       c.Status(404)
//        return c.JSON(fiber.Map{
//         "message": "Invalid Token",
//       })
      
      
//     }
//     timeCreated:=resetpassword.CreatedAt
//     expiredTime:=timeCreated.Add(2 * time.Hour)
//     compareDate:=time.Now().After(expiredTime)
//     if compareDate{
//       //database.DB.Where("user_id = ?", user.Id).Delete(&confirmCode)
//       c.Status(404)
//       return c.JSON(fiber.Map{
//        "message": "Password reset code is expired. ",
//      })

//     }
//     database.DB.Model(&resetpassword).Where("token = ?", data["token"]).Update("used", 1)
//     user:=c.Locals("user").(*jwt.Token)
// 	claims:=user.Claims.(jwt.MapClaims)
// 	id:=claims["user_id"].(float64)

//     user:= model.User{}
//     //Check if the password length is more than 6
//   if len(data["password"])  <= 6 {
//     c.Status(400)
//     return c.JSON(fiber.Map{
//     "message": "Password must be more than 6 character",
//   })
     
//   }
//   hash, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
//   if err != nil {
// 	  return err
//   }
//     database.DB.Model(&user).Where("id=?", id).Updates(user)
//     database.DB.Where("used = ?", 1).Delete(&resetpassword)

//     return c.JSON(fiber.Map{
//       "message": "Password reset successfully ",
//     })


//   }

// func Home(c *fiber.Ctx) error  {
// 	   return c.JSON(fiber.Map{"success": true, "path": "public"})
	
	
// }

// func Private(c *fiber.Ctx) error { 
// 	user:=c.Locals("user").(*jwt.Token)
// 	claims:=user.Claims.(jwt.MapClaims)
// 	id:=claims["user_id"].(float64)


// 	return c.JSON(fiber.Map{"success": true, "path": "private","id":id })
	
	
// }
