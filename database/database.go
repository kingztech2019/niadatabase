package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kingztech2019/nia_backend/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB
func Connect()  {
	err:=godotenv.Load()
	if err != nil {
		 log.Fatal("Error load .env file")
	}

	
	dbUser:=os.Getenv("DB_USER")
	dbPass:=os.Getenv("DB_PASSWORD")
	dbHost:=os.Getenv("DB_HOST")
	dbName:=os.Getenv("DB_NAME")
	dbPort:=os.Getenv("DB_PORT")
 
	dsn:= fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",dbHost,dbUser,dbPass,dbName,dbPort)
	database,err:=gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		panic("Could not connect to the database")
		
	}else{
		log.Println("connect successfully")
	} 
	DB=database
	database.AutoMigrate(
		&model.User{},
		&model.VechicleInsurance{},
		&model.VerifyCode{},
		&model.ImagesUrl{},
		&model.PersonalDetails{},
		&model.PolicyDetails{},
		 
	)

	
}