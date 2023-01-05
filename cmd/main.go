package main

import (
	"github.com/bootcamp/supermercadito/cmd/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	handlers.SetProductos()
	handlers.SetProductGroupRoutes(server)

	server.Run(":8080")
}
