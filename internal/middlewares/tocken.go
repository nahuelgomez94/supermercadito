package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func TokenValidation() gin.HandlerFunc {
	TOKEN := os.Getenv("TOKEN")

	return func(c *gin.Context) {
		hToken := c.GetHeader("token")

		if hToken != TOKEN {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid API Token"})
			return
		}

		c.Next()
	}
}
