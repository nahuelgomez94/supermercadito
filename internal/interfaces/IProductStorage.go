package interfaces

import "github.com/bootcamp/supermercadito/internal/dto"

type IProductStorage interface {
	Get() ([]dto.Producto, error)
	Set(products []dto.Producto) (err error)
}
