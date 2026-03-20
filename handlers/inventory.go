package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ibraah007/hospital-inventory-api/models"
)

// The "Shelf" where data is stored in memory
var inventory = []models.Item{
	{ID: "1", Name: "Paracetamol", Quantity: 500},
	{ID: "2", Name: "Surgical Masks", Quantity: 2000},
}

// GetInventory handles GET /inventory (View all)
func GetInventory(c *gin.Context) {
	c.JSON(http.StatusOK, inventory)
}

// AddItem handles POST /inventory (Add new)
func AddItem(c *gin.Context) {
	var newItem models.Item

	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inventory = append(inventory, newItem)
	c.JSON(http.StatusCreated, newItem)
}

// DeleteItem handles DELETE /inventory/:id (Remove one)
func DeleteItem(c *gin.Context) {
	id := c.Param("id") // Grab the ID from the URL

	for i, item := range inventory {
		if item.ID == id {
			// The "Snip": Join everything before 'i' and everything after 'i'
			inventory = append(inventory[:i], inventory[i+1:]...)
			
			c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
			return
		}
	}

	// If we looped through everything and didn't find the ID
	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}