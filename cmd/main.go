package main

import (
	"github.com/bootcamp/supermercadito/cmd/handlers"
	"github.com/bootcamp/supermercadito/internal/producto"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	producto.InicializarDB()

	// Armo los endpoint del server
	SetProductGroupRoutes(server)

	server.Run(":8080")
}

func SetProductGroupRoutes(server *gin.Engine) {
	p := server.Group("/products")

	p.GET("/", handlers.GetProductos)
	p.GET("/:id", handlers.GetProductoById)
	p.GET("/search", handlers.GetProductsByMinPrice)
	p.POST("/", handlers.SetProducto)
}
