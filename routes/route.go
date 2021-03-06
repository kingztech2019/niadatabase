package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/kingztech2019/nia_backend/controller"
)


func Setup(app *fiber.App)  {
	private := app.Group("/private")
	private.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	private.Post("/buy-details", controller.VechicleInsurance)
	private.Post("/all-details", controller.PolicyDetails)
	private.Post("/personal-details", controller.PersonalDetails)
	private.Get("/status-check", controller.GetStatus)
	private.Post("/get-vin-details", controller.CheckVin)
	// private.Get("/", controller.Private )
	
	public := app.Group("/public")
	// public.Get("/",controller.Home)
	public.Post("/signup", controller.SignUp)
	public.Post("/uploadimage", controller.UploadImage)
	public.Static("/uploads", "./uploads")
	public.Post("/login", controller.Login)
	public.Post("/password-reset-code", controller.CheckEmailPaswordReset)
	public.Post("/reset-password", controller.ChangePassword)
	

}
