package producto

import (
	"github.com/bootcamp/supermercadito/internal/dto"
)

var pr productRepository

func SetProductos() {
	pr.SetProductos()
}

func GetProductos() []dto.Producto {
	return pr.getProductos()
}

func GetProductoById(id int) dto.Producto {
	return pr.getProductoById(id)
}

func GetProductsByMinPrice(price float64) []dto.Producto {
	return pr.getProductsByMinPrice(price)
}

func SetProducto(newProduct dto.Producto) (savedProd dto.Producto, err error) {
	savedProd, err = pr.SetProducto(newProduct)
	return
}
