package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ibraah007/hospital-inventory-api/handlers"
)

func main() {
	r := gin.Default()

	// The Manager (Gin) now points to the Specialists (Handlers)
	r.GET("/inventory", handlers.GetInventory)
	r.POST("/inventory", handlers.AddItem)

	r.Run(":8080")
}