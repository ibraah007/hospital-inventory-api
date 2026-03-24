package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ibraah007/hospital-inventory-api/database"
	"github.com/ibraah007/hospital-inventory-api/models"
)

func GetInventory(c *gin.Context) {
	var items []models.Inventory
	database.DB.Find(&items)
	c.JSON(http.StatusOK, items)
}

func SearchInventory(c *gin.Context) {
	query := c.Query("name")
	var items []models.Inventory
	// Search by name OR department
	database.DB.Where("name LIKE ? OR department LIKE ?", "%"+query+"%", "%"+query+"%").Find(&items)
	c.JSON(http.StatusOK, items)
}

func AddItem(c *gin.Context) {
	var input models.Inventory
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&input)
	c.JSON(http.StatusOK, input)
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Inventory{}, "id = ?", id)
	c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
}