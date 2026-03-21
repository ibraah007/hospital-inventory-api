package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ibraah007/hospital-inventory-api/database"
	"github.com/ibraah007/hospital-inventory-api/handlers"
	"github.com/ibraah007/hospital-inventory-api/middleware"
)

func main() {
	database.InitDB()
	r := gin.Default()

	// CORS Setup
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-API-KEY")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// PUBLIC
	r.GET("/staff", handlers.GetStaff)
	r.GET("/staff/search", handlers.SearchStaff)
	r.GET("/inventory", handlers.GetInventory)
	r.GET("/inventory/search", handlers.SearchInventory)

	// PROTECTED
	protected := r.Group("/")
	protected.Use(middleware.AuthGuard()) 
	{
		protected.POST("/staff", handlers.AddStaff)
		protected.DELETE("/staff/:id", handlers.DeleteStaff) 
		protected.POST("/inventory", handlers.AddItem)
		protected.DELETE("/inventory/:id", handlers.DeleteItem)
	}

	r.Run(":8080")
}