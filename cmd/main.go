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

	// Armo los endpoint del server
	SetProductGroupRoutes(server)

	server.Run(":8080")
}

func SetProductGroupRoutes(server *gin.Engine) {

	var productHandler handlers.ProductHandler

	repo := producto.NewProductRepository()
	productHandler.ProductService = *producto.NewProductService(repo)

	p := server.Group("/products")
	p.GET("/", productHandler.GetProductos)
	p.GET("/:id", productHandler.GetProductoById)
	p.GET("/search", productHandler.GetProductsByMinPrice)
	p.POST("/", productHandler.SetProducto)
}
