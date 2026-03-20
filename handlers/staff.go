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

// SearchStaff handles GET /staff/search?name=...
func SearchStaff(c *gin.Context) {
	name := c.Query("name") // This grabs the name from the URL
	var staff []models.Staff

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide a name to search"})
		return
	}

	// Search the database where name LIKE the input
	// The "%" symbols allow for partial matches (e.g., "Ibra" finds "Ibrahim")
	result := database.DB.Where("name LIKE ?", "%"+name+"%").Find(&staff)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No doctor found with that name"})
		return
	}

	c.JSON(http.StatusOK, staff)
}