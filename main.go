package main

import (
	"log"
	"math/rand"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/kingztech2019/nia_backend/database"
	"github.com/kingztech2019/nia_backend/routes"
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
	err:=godotenv.Load()
	if err != nil {
		 log.Fatal("Error loading .env files")
	}
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
}))
 
	routes.Setup(app)
	port:=os.Getenv("PORT")


	if err := app.Listen(":"+port); err != nil {
		panic(err)
	}
	
}


