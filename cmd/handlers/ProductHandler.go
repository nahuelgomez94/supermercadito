package handlers

import (
	"net/http"
	"strconv"

	"github.com/bootcamp/supermercadito/internal/dto"
	"github.com/bootcamp/supermercadito/internal/producto"
	"github.com/bootcamp/supermercadito/internal/web"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductService producto.ProductService
}

const (
	ErrIdFormatInvalid = "ID con formato inválido"
	ErrNoFoundId       = "No se ha encontrado el ID especificado"
)

func NewProductHandler(ps producto.ProductService) (ph *ProductHandler) {
	return &ProductHandler{
		ProductService: ps,
	}
}

// ListProducts godoc
// @Summary List all products
// @Tags Products
// @Description Get All Products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200
// @Router /products [get]
func (ph *ProductHandler) GetProductos(c *gin.Context) {
	rta, err := ph.ProductService.GetProductos()
	if err != nil {
		web.NewErrorResponse(c, http.StatusBadRequest, "", err.Error())
		return
	}

	web.NewResponse(c, http.StatusOK, rta)
}

func (ph *ProductHandler) GetProductoById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		web.NewErrorResponse(c, http.StatusBadRequest, "BadRequest", ErrIdFormatInvalid)
		return
	}

	rta, err := ph.ProductService.GetProductoById(id)

	if err != nil {
		web.NewErrorResponse(c, http.StatusInternalServerError, "InternalServerError", err.Error())
		return
	}

	if rta.Id != id {
		web.NewErrorResponse(c, http.StatusNotFound, "NotFound", ErrNoFoundId)
		return
	}

	web.NewResponse(c, http.StatusOK, rta)
}

func (ph *ProductHandler) SetProducto(c *gin.Context) {

	var prodReq dto.ProductoRequest
	err := c.ShouldBindJSON(&prodReq)

	if err != nil {
		web.NewErrorResponse(c, http.StatusBadRequest, "BadRequest", err.Error())
		return
	}

	var prodConvert dto.Producto
	prodConvert.SetValues(prodReq)
	savedProd, err := ph.ProductService.SetProducto(dto.Producto(prodConvert))

	if err != nil {
		web.NewErrorResponse(c, http.StatusBadRequest, "BadRequest", err.Error())
		return
	}

	web.NewResponse(c, http.StatusOK, savedProd)
}

func (ph *ProductHandler) GetProductsByMinPrice(c *gin.Context) {

	pPrice := c.Query("pricegt")
	price, err := strconv.ParseFloat(pPrice, 10)

	if err != nil {
		web.NewErrorResponse(c, http.StatusBadRequest, "BadRequest", err.Error())
		return
	}

	rta, err := ph.ProductService.GetProductsByMinPrice(price)

	if err != nil {
		web.NewErrorResponse(c, http.StatusBadRequest, "BadRequest", err.Error())
		return
	}
	web.NewResponse(c, http.StatusOK, rta)
}

func (ph *ProductHandler) UpdateProduct(c *gin.Context) {
	var prodReq dto.ProductoRequest
	err := c.ShouldBindJSON(&prodReq)

	if err != nil {
		web.NewErrorResponse(c, http.StatusBadRequest, "BadRequest", err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		web.NewErrorResponse(c, http.StatusBadRequest, "BadRequest", ErrIdFormatInvalid)
		return
	}

	prod, err := ph.ProductService.UpdateProduct(id, prodReq)

	if err != nil {
		web.NewErrorResponse(c, http.StatusInternalServerError, "InternalServerError", "")
		return
	}

	web.NewResponse(c, http.StatusOK, prod)
}

func (ph *ProductHandler) PatchProduct(c *gin.Context) {

	var cambios dto.Producto
	err := c.ShouldBindJSON(&cambios)

	if err != nil {
		web.NewErrorResponse(c, http.StatusBadRequest, "BadRequest", err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		web.NewErrorResponse(c, http.StatusBadRequest, "BadRequest", ErrIdFormatInvalid)
		return
	}

	prodPatcheado, err := ph.ProductService.PatchProduct(id, cambios)

	if err != nil {
		web.NewErrorResponse(c, http.StatusInternalServerError, "InternalServerError", "")
		return
	}

	web.NewResponse(c, http.StatusOK, prodPatcheado)
}

func (ph *ProductHandler) DeleteProduct(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		web.NewErrorResponse(c, http.StatusBadRequest, "BadRequest", ErrIdFormatInvalid)
		return
	}

	err = ph.ProductService.DeleteProduct(id)

	if err != nil {
		web.NewErrorResponse(c, http.StatusInternalServerError, "InternalServerError", "")
		return
	}

	web.NewResponse(c, http.StatusNoContent, nil)

}
