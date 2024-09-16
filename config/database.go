package config

import (
	"github.com/glinboy/gin-oauth2-secured-demo/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DatabaseConnecting() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&model.Book{})
	if err != nil {
		panic(err)
	}

	return db
}
