package routers

import (
	"github.com/bootcamp/supermercadito/cmd/handlers"
	"github.com/bootcamp/supermercadito/internal/producto"
	"github.com/gin-gonic/gin"
)

type Router struct {
	en *gin.Engine
	db *producto.ProductRepository
}

func NewRouter(en *gin.Engine, db *producto.ProductRepository) *Router {
	return &Router{en: en, db: db}
}

func (r *Router) SetProductGroupRoutes(server *gin.Engine) {
	service := producto.NewProductService(r.db)
	productHandler := handlers.NewProductHandler(*service)

	p := server.Group("/products")
	p.GET("/", productHandler.GetProductos)
	p.GET("/:id", productHandler.GetProductoById)
	p.GET("/search", productHandler.GetProductsByMinPrice)
	p.POST("/", productHandler.SetProducto)
	p.PUT("/:id", productHandler.UpdateProduct)
	p.PATCH("/:id", productHandler.PatchProduct)
}
