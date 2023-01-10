package main

import (
	"github.com/bootcamp/supermercadito/internal/producto"
	"github.com/bootcamp/supermercadito/internal/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Armo los endpoint del server
	repo := producto.NewProductRepository()
	router := routers.NewRouter(server, repo)
	router.SetProductGroupRoutes(server)

	server.Run(":8080")
}
