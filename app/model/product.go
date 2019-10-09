package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&User{})
	return db
}
