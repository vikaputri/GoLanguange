package config

import (
	"restApi/structs"

	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:admin@tcp(localhost:3306)/godb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failde to connect to database")
	}

	db.AutoMigrate(structs.Person{})
	return db
}
