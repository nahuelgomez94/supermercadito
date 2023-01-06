package handlers

import (
	"strconv"

	"github.com/bootcamp/supermercadito/internal/dto"
	"github.com/bootcamp/supermercadito/internal/producto"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	// Interfaz?
	ProductService producto.ProductService
}

func (ph *ProductHandler) GetProductos(c *gin.Context) {
	rta, err := ph.ProductService.GetProductos()
	if err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, rta)
}

func (ph *ProductHandler) GetProductoById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.String(500, "ID con formato inválido")
		return
	}

	rta, err := ph.ProductService.GetProductoById(id)

	if err != nil {
		c.JSON(500, err.Error())
	}

	if rta.Id != id {
		c.String(500, "No se ha encontrado el ID especificado")
		return
	}

	c.JSON(200, rta)
}

func (ph *ProductHandler) SetProducto(c *gin.Context) {
	var prodReq dto.ProductoRequest
	err := c.ShouldBindJSON(&prodReq)

	if err != nil {
		c.String(401, err.Error())
		return
	}

	savedProd, err := ph.ProductService.SetProducto(dto.Producto(prodReq))

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, savedProd)
}

func (ph *ProductHandler) GetProductsByMinPrice(c *gin.Context) {
	pPrice := c.Query("pricegt")
	price, err := strconv.ParseFloat(pPrice, 10)

	if err != nil {
		c.String(500, err.Error())
		return
	}

	rta, err := ph.ProductService.GetProductsByMinPrice(price)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	if len(rta) == 0 {
		c.String(200, "No hay productos con ese precio")
		return
	}

	c.JSON(200, rta)
}
