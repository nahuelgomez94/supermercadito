package handlers

import (
	"strconv"

	"github.com/bootcamp/supermercadito/internal/dto"
	"github.com/bootcamp/supermercadito/internal/producto"
	"github.com/gin-gonic/gin"
)

func SetProductGroupRoutes(server *gin.Engine) {
	producto.SetProductos()
	p := server.Group("/products")

	p.GET("/", func(c *gin.Context) {
		c.JSON(200, producto.GetProductos())
	})

	p.GET("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.String(500, "Error en la conversion ID")
			return
		}
		if id <= 0 {
			c.String(500, "No se ha especificado ID")
			return
		}

		rta := producto.GetProductoById(id)

		if rta.Id != id {
			c.String(500, "No se ha encontrado el ID especificado")
			return
		}

		c.JSON(200, rta)
	})

	p.GET("/search", func(c *gin.Context) {
		pPrice := c.Query("pricegt")
		price, err := strconv.ParseFloat(pPrice, 10)

		if err != nil {
			c.String(500, err.Error())
			return
		}

		if price < 0 {
			c.String(500, "El precio no puede ser negativo")
			return
		}

		rta := producto.GetProductsByMinPrice(price)

		if len(rta) == 0 {
			c.String(200, "No hay productos con ese precio")
			return
		}

		c.JSON(200, rta)

	})

	p.POST("/", func(c *gin.Context) {
		var prodReq dto.Producto
		err := c.ShouldBindJSON(&prodReq)

		if err != nil {
			c.String(401, err.Error())
			return
		}

		savedProd, err := producto.SetProducto(dto.Producto(prodReq))

		if err != nil {
			c.JSON(500, "Internal Server Error")
		}

		c.JSON(200, savedProd)
	})
}

func SetProductos() (err error) {
	producto.SetProductos()

	return nil
}
