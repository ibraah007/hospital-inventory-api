package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ibraah007/hospital-inventory-api/database"
	"github.com/ibraah007/hospital-inventory-api/models"
)

// GetStaff handles GET /staff (Fetches everyone from the DB)
func GetStaff(c *gin.Context) {
	var staff []models.Staff
	
	// .Find() is a GORM command that says: "Get everything in the Staff table"
	if err := database.DB.Find(&staff).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch staff"})
		return
	}
	
	c.JSON(http.StatusOK, staff)
}

// AddStaff handles POST /staff (Saves a new person to the DB)
func AddStaff(c *gin.Context) {
	var newStaff models.Staff

	if err := c.ShouldBindJSON(&newStaff); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// .Create() is a GORM command that saves the data to your hard drive
	database.DB.Create(&newStaff)
	c.JSON(http.StatusCreated, newStaff)
}