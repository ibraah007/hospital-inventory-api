package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ibraah007/hospital-inventory-api/models"
)

// The "Shelf" - Current stock in memory
var inventory = []models.Item{
	{ID: "1", Name: "Paracetamol", Quantity: 500},
	{ID: "2", Name: "Surgical Masks", Quantity: 2000},
}

// GetInventory handles GET /inventory (Read All)
func GetInventory(c *gin.Context) {
	c.JSON(http.StatusOK, inventory)
}

// AddItem handles POST /inventory (Create)
func AddItem(c *gin.Context) {
	var newItem models.Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	inventory = append(inventory, newItem)
	c.JSON(http.StatusCreated, newItem)
}

// UpdateItem handles PUT /inventory/:id (Update)
func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var updatedData models.Item

	// Validate the incoming JSON
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the item and swap it with the new data
	for i, item := range inventory {
		if item.ID == id {
			inventory[i] = updatedData
			c.JSON(http.StatusOK, inventory[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

// DeleteItem handles DELETE /inventory/:id (Delete)
func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	for i, item := range inventory {
		if item.ID == id {
			inventory = append(inventory[:i], inventory[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}