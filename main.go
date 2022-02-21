package main

import (
	"math/rand"

	"github.com/floydjones1/auth-server/database"
	"github.com/floydjones1/auth-server/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)
 
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}


func main() {
	
	database.Connect()
 
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
}))
 
	routes.Setup(app)
	


	if err := app.Listen(":4000"); err != nil {
		panic(err)
	}
	
}


