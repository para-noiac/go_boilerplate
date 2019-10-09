package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// type Model struct {
// 	DB *gorm.DB
// }

// func (model *Model) Modelling() {
// 	model.DB = DBMigrate()
// }

func DBMigrate(db *gorm.DB) *gorm.DB {
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
