package routes

import (
	"github.com/floydjones1/auth-server/controller"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)


func Setup(app *fiber.App)  {
	private := app.Group("/private")
	private.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	private.Post("/buy-details", controller.VechicleInsurance)
	private.Post("/all-details", controller.PolicyDetails)
	private.Post("/personal-details", controller.PersonalDetails)
	// private.Get("/", controller.Private )
	
	public := app.Group("/public")
	// public.Get("/",controller.Home)
	public.Post("/signup", controller.SignUp)
	public.Post("/uploadimage", controller.UploadImage)
	public.Static("/uploads", "./uploads")
	public.Post("/login", controller.Login)
}
