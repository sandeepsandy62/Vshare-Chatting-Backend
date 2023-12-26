package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDB initializes the database connection using GORM
func InitDB(dbHost string, dbUser string, dbPassword string, dbName string, dbPort string) error {


	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database : ", err)
		return err
	}

	return nil

}

//return DB instance
func GetDB() *gorm.DB {
	return db
}
