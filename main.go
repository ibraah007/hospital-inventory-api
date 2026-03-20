package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ibraah007/hospital-inventory-api/handlers"
)

func main() {
	r := gin.Default()

	// Route 1: See everything
	r.GET("/inventory", handlers.GetInventory)
	
	// Route 2: Add something new
	r.POST("/inventory", handlers.AddItem)
	
	// Route 3: Delete something specific using its ID
	r.DELETE("/inventory/:id", handlers.DeleteItem)

	// Start the server on port 8080
	r.Run(":8080")
}