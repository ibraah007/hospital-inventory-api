func main() {
    database.InitDB() // This opens the filing cabinet

    r := gin.Default()

    // Pharmacy Department
    r.GET("/inventory", handlers.GetInventory)
    r.POST("/inventory", handlers.AddItem)

    // HR Department (Make sure these two lines are there!)
    r.GET("/staff", handlers.GetStaff)
    r.POST("/staff", handlers.AddStaff)

    r.Run(":8080")
}