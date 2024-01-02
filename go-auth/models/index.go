package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

var DB *gorm.DB

func InitDB(cfg Config) {

	//data source name [dsn]
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	//Postgres.Open call add dsn to config object
	//config object also contains some extra fields
	//then convert the config object to dialector and pass it to the gorm.Open
	//gorm.Open returns db pointer and error
	db, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic(err)
	}

	/*
		AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes.
		It will change existing column's type if its size, precision, nullable changed.
		It WON'T delete unused columns to protect your data.
	*/
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}

	fmt.Println("Migrated database")

	DB = db

}
