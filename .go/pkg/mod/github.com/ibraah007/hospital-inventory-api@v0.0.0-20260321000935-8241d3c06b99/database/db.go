package database

import (
	"github.com/ibraah007/hospital-inventory-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("hospital.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	// AutoMigrate creates/updates tables based on your models
	DB.AutoMigrate(&models.Staff{}, &models.Inventory{})
}