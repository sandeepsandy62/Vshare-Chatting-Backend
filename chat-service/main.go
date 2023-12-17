package main

import (
	"gorm.io/gorm"
	"log"
	"os"
	

	"vshare.com/common/database"
	"github.com/joho/godotenv"
)

func main() {

	envFilePath := "../.env"
	if err := godotenv.Load(envFilePath); err != nil {
		log.Fatal("Error loading the .env file %v ", err)
	} else {
		log.Print(".env file loaded successfully!")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")




	var err error = database.InitDB(dbHost,dbUser,dbPassword,dbName,dbPort)
	if err != nil {
		log.Fatalf("Error connecting to DB %v", err)
	} else {
		var db *gorm.DB = database.GetDB()
		log.Print("Connected to DB : %v", db)
	}
}
