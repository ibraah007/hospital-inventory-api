package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ibraah007/hospital-inventory-api/database"
	"github.com/ibraah007/hospital-inventory-api/handlers"
	"github.com/ibraah007/hospital-inventory-api/middleware"
)

func main() {
	// 1. Initialize the Database (The Filing Cabinet)
	database.InitDB()

	// 2. Initialize the Router (The Receptionist)
	r := gin.Default()

	// --- PUBLIC ROUTES (Anyone can access these) ---
	
	// Inventory View
	r.GET("/inventory", handlers.GetInventory)
	
	// Staff View & Search
	r.GET("/staff", handlers.GetStaff)
	r.GET("/staff/search", handlers.SearchStaff) // <--- THIS fixes your 404!

	// --- PROTECTED ROUTES (Requires X-API-KEY: HospitalSecret123) ---
	
	protected := r.Group("/")
	protected.Use(middleware.AuthGuard()) 
	{
		// Inventory Management
		protected.POST("/inventory", handlers.AddItem)
		protected.PUT("/inventory/:id", handlers.UpdateItem)
		protected.DELETE("/inventory/:id", handlers.DeleteItem)

		// Staff Management (Hiring)
		protected.POST("/staff", handlers.AddStaff)
	}

	// 3. Start the Server
	r.Run(":8080")
}