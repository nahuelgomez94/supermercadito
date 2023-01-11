package main

import (
	"github.com/bootcamp/supermercadito/internal/producto"
	"github.com/bootcamp/supermercadito/internal/routers"
	"github.com/bootcamp/supermercadito/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Armo los endpoint del server
	storage := store.NewProductStorage()
	repo := producto.NewProductRepository(storage)
	router := routers.NewRouter(server, repo)
	router.SetProductGroupRoutes(server)

	server.Run(":8080")
}
