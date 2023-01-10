package handlers

import (
	"strconv"

	"github.com/bootcamp/supermercadito/internal/dto"
	"github.com/bootcamp/supermercadito/internal/producto"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductService producto.ProductService
}

func NewProductHandler(ps producto.ProductService) (ph *ProductHandler) {
	return &ProductHandler{
		ProductService: ps,
	}
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
		return
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

	var prodConvert dto.Producto
	prodConvert.SetValues(prodReq)
	savedProd, err := ph.ProductService.SetProducto(dto.Producto(prodConvert))

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

func (ph *ProductHandler) UpdateProduct(c *gin.Context) {
	var prodReq dto.ProductoRequest
	err := c.ShouldBindJSON(&prodReq)

	if err != nil {
		c.String(500, "Solicitud inválida")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.String(500, "Parametro ID inválido")
		return
	}

	prod, err := ph.ProductService.UpdateProduct(id, prodReq)

	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.JSON(200, prod)
}

func (ph *ProductHandler) PatchProduct(c *gin.Context) {
	var cambios dto.Producto
	err := c.ShouldBindJSON(&cambios)

	if err != nil {
		c.String(500, "Error al interpretar el objeto")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.String(500, "Id Inválido")
		return
	}

	prodPatcheado, err := ph.ProductService.PatchProduct(id, cambios)

	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.JSON(200, prodPatcheado)
}
