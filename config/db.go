package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("Host"),
		os.Getenv("Port"),
		os.Getenv("User"),
		os.Getenv("Password"),
		os.Getenv("DBname"),
	)

	db, err := gorm.Open(os.Getenv("Dialect"), dbURI)
	if err != nil {
		panic(err)
	}
	return db
}
