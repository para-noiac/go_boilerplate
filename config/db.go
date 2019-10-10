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
		os.Getenv("DB_Host"),
		os.Getenv("DB_Port"),
		os.Getenv("DB_User"),
		os.Getenv("DB_Password"),
		os.Getenv("DB_Name"),
	)

	db, err := gorm.Open(os.Getenv("DB_Dialect"), dbURI)
	if err != nil {
		panic(err)
	}
	return db
}
