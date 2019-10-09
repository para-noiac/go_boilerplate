package model

import (
	"boilerplate/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func DBMigrate() *gorm.DB {
	db := config.InitDB()
	/*
		----------------------------------------
		| Auto migration to ease table changes |
		----------------------------------------
		Add struct name from model folders to run auto migration (e.g. &Product{})
		More functions refer to http://jinzhu.me/gorm/database.html#migration
	*/
	db.AutoMigrate(&Product{}, &User{})
	return db
}
