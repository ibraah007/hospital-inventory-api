package database

import (
	"github.com/ibraah007/hospital-inventory-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	// This creates a file named 'hospital.db' in your project folder
	DB, err = gorm.Open(sqlite.Open("hospital.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// This automatically creates the "Items" and "Staff" tables for you!
	DB.AutoMigrate(&models.Item{}, &models.Staff{})
}