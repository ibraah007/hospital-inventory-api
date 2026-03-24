package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ibraah007/hospital-inventory-api/database"
	"github.com/ibraah007/hospital-inventory-api/models"
)

func GetStaff(c *gin.Context) {
	var staff []models.Staff
	database.DB.Find(&staff)
	c.JSON(http.StatusOK, staff)
}

func SearchStaff(c *gin.Context) {
	query := c.Query("name")
	var staff []models.Staff
	database.DB.Where("name LIKE ? OR role LIKE ? OR shift LIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%").Find(&staff)
	c.JSON(http.StatusOK, staff)
}

func AddStaff(c *gin.Context) {
	var input models.Staff
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&input)
	c.JSON(http.StatusOK, input)
}

func DeleteStaff(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Staff{}, "id = ?", id)
	c.JSON(http.StatusOK, gin.H{"message": "Staff removed"})
}