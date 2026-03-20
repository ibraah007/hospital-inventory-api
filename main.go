package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ibraah007/hospital-inventory-api/handlers"
)

func main() {
	// Initialize the Manager (Gin)
	r := gin.Default()

	// The Switchboard: URL + Method -> Specialist Function
	r.GET("/inventory", handlers.GetInventory)      // See all
	r.POST("/inventory", handlers.AddItem)          // Add new
	r.PUT("/inventory/:id", handlers.UpdateItem)    // Change existing
	r.DELETE("/inventory/:id", handlers.DeleteItem) // Remove

	// Start the server
	r.Run(":8080")
}