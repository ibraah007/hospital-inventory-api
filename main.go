package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	
	
	"github.com/ibraah007/hospital-inventory-api/models" 
)

/
var inventory = []models.Item{
	{ID: "1", Name: "Paracetamol", Quantity: 500},
	{ID: "2", Name: "Surgical Masks", Quantity: 2000},
}

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "Hospital API is Online"})
	})

	r.GET("/inventory", func(c *gin.Context) {
		c.JSON(http.StatusOK, inventory)
	})

	r.Run(":8080")
}