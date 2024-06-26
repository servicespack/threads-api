package configs

import (
	"github.com/servicespack/threads-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitializeDatabase() {
	database, err := gorm.Open(sqlite.Open("threads.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db = database
	db.AutoMigrate(&models.Thread{})
}

func GetDB() *gorm.DB {
	return db
}
