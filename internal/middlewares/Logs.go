package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func NewLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		hhmmss := time.Now().UTC().Format("2006-01-02 15:04:05")
		c.Next()
		fmt.Printf("Hola! Soy un %s - SON LAS %s - Mi url es: %s - Peso %d bytes\n", c.Request.Method, hhmmss, c.Request.RequestURI, c.Request.ContentLength)
	}
}
