package configs

import (
	"github.com/servicespack/threads-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeDatabase() {
	database, err := gorm.Open(sqlite.Open("threads.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&models.Thread{})
}
