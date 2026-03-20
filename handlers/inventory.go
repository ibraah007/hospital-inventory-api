package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ibraah007/hospital-inventory-api/database"
	"github.com/ibraah007/hospital-inventory-api/models"
)

// GetInventory fetches everything from the SQLITE database
func GetInventory(c *gin.Context) {
	var items []models.Item
	database.DB.Find(&items)
	c.JSON(http.StatusOK, items)
}

// AddItem saves a new item to the SQLITE database
func AddItem(c *gin.Context) {
	var newItem models.Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&newItem)
	c.JSON(http.StatusCreated, newItem)
}

// UpdateItem updates an item in the DB by ID
func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item

	if err := database.DB.First(&item, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&item)
	c.JSON(http.StatusOK, item)
}

// DeleteItem removes an item from the DB by ID
func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Item{}, "id = ?", id)
	c.JSON(http.StatusOK, gin.H{"message": "Item deleted from database"})
}