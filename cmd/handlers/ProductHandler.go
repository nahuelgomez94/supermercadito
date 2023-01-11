package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
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

func (ph *ProductHandler) TokenValidation(c *gin.Context) (auth bool, err error) {
	TOKEN := os.Getenv("TOKEN")

	hToken := c.GetHeader("token")

	if hToken == TOKEN {
		return true, nil
	} else {
		return false, errors.New("No auth")
	}
}

func (ph *ProductHandler) GetProductos(c *gin.Context) {
	auth, errorAuth := ph.TokenValidation(c)

	if errorAuth != nil && auth == false {
		c.JSON(http.StatusUnauthorized, errorAuth)
		return
	}

	rta, err := ph.ProductService.GetProductos()
	if err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, rta)
}

func (ph *ProductHandler) GetProductoById(c *gin.Context) {
	auth, errorAuth := ph.TokenValidation(c)

	if errorAuth != nil && auth == false {
		c.JSON(http.StatusUnauthorized, errorAuth)
		return
	}

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
	auth, errorAuth := ph.TokenValidation(c)

	if errorAuth != nil && auth == false {
		c.JSON(http.StatusUnauthorized, errorAuth)
		return
	}

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
	auth, errorAuth := ph.TokenValidation(c)

	if errorAuth != nil && auth == false {
		c.JSON(http.StatusUnauthorized, errorAuth)
		return
	}

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
	auth, errorAuth := ph.TokenValidation(c)

	if errorAuth != nil && auth == false {
		c.JSON(http.StatusUnauthorized, errorAuth)
		return
	}

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
	auth, errorAuth := ph.TokenValidation(c)

	if errorAuth != nil && auth == false {
		c.JSON(http.StatusUnauthorized, errorAuth)
		return
	}

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

func (ph *ProductHandler) DeleteProduct(c *gin.Context) {
	auth, errorAuth := ph.TokenValidation(c)

	if errorAuth != nil && auth == false {
		c.JSON(http.StatusUnauthorized, errorAuth)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	err = ph.ProductService.DeleteProduct(id)

	if err != nil {
		fmt.Println(err.Error())
		c.String(http.StatusNotModified, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
