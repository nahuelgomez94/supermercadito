package interfaces

import "github.com/bootcamp/supermercadito/internal/dto"

type IProductService interface {
	GetProductos() (rta []dto.Producto, err error)
	GetProductsByMinPrice(price float64) (rta []dto.Producto, err error)
	GetProductoById(id int) (rta dto.Producto, err error)
	SetProducto(newProduct dto.Producto) (savedProd dto.Producto, err error)
	validateFormatDate(date string) (err error)
}
