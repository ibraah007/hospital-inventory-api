package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// This is the "Password" for your hospital system
const MasterKey = "HospitalSecret123" 

func AuthGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		// We look for a "Badge" in the request header called X-API-KEY
		apiKey := c.GetHeader("X-API-KEY")

		if apiKey != MasterKey {
			// If the key is wrong or missing, Slam the door!
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Access Denied"})
			c.Abort() 
			return
		}

		// If the key is correct, let them in
		c.Next()
	}
}