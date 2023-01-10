package interfaces

import "github.com/bootcamp/supermercadito/internal/dto"

type IProductRepository interface {
	SetProductos() (err error)
	GetProductos() (rta []dto.Producto, err error)
	GetProductoById(id int) (rta dto.Producto, err error)
	GetProductsByMinPrice(price float64) (rta []dto.Producto, err error)
	ValidateExistsCodeProduct(code string) (err error)
	SetProducto(newProduct dto.Producto) (savedProd dto.Producto, err error)
	UpdateProduct(id int, producto dto.Producto) (savedProd dto.Producto, err error)
	ValidateUniqueCode(id int, code string) (err error)
}
