package main

import (
	"go-auth/models"
	"go-auth/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	//create a new gin instance
	r := gin.Default()

	//load .env file and create a new connectiong to the database
	envFilePath := "../.env"
	if err := godotenv.Load(envFilePath); err != nil {
		log.Fatal("Error loading the .env file %v ", err)
	} else {
		log.Print(".env file loaded successfully!")
	}

	config := models.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	//initilize DB
	models.InitDB(config)

	//Load the routes
	routes.AuthRoutes(r)

	r.Run(":8080")

}
