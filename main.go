package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/ibraah007/hospital-inventory-api/models" 
)

var inventory = []models.Item{
	{ID: "1", Name: "Paracetamol", Quantity: 500},
	{ID: "2", Name: "Surgical Masks", Quantity: 2000},
}

func main() {
	r := gin.Default()

	// 1. View all items
	r.GET("/inventory", func(c *gin.Context) {
		c.JSON(http.StatusOK, inventory)
	})

	// 2. Add a new item (The Loading Dock)
	r.POST("/inventory", func(c *gin.Context) {
		var newItem models.Item

		// Call BindJSON to bind the received JSON to newItem
		if err := c.ShouldBindJSON(&newItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Add the new item to the slice
		inventory = append(inventory, newItem)
		c.JSON(http.StatusCreated, newItem)
	})

	r.Run(":8080")
}