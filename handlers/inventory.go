package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ibraah007/hospital-inventory-api/database"
	"github.com/ibraah007/hospital-inventory-api/models"
	"net/http"
)

func GetInventory(c *gin.Context) {
	var items []models.Item
	database.DB.Find(&items)
	c.JSON(http.StatusOK, items)
}

func SearchInventory(c *gin.Context) {
	name := c.Query("name")
	var items []models.Item
	database.DB.Where("name LIKE ?", "%"+name+"%").Find(&items)
	c.JSON(http.StatusOK, items)
}

func AddItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Data"})
		return
	}
	database.DB.Save(&item)
	c.JSON(http.StatusCreated, item)
}

func RefillItem(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		Quantity int `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity"})
		return
	}

	var item models.Item
	if err := database.DB.First(&item, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	item.Quantity += input.Quantity
	database.DB.Save(&item)
	c.JSON(http.StatusOK, item)
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Item{}, "id = ?", id)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}