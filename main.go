package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// This uses the "net/http" package (http.StatusOK)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Hospital API is Online",
		})
	})

	r.Run(":8080")
}