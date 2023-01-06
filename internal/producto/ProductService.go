package producto

import (
	"errors"
	"time"

	"github.com/bootcamp/supermercadito/internal/dto"
	"github.com/bootcamp/supermercadito/internal/interfaces"
)

type ProductService struct {
	interfaces.IProductService
	pr interfaces.IProductRepository
}

func NewProductService(repo interfaces.IProductRepository) *ProductService {
	return &ProductService{
		pr: repo,
	}
}

func (p ProductService) GetProductos() (rta []dto.Producto, err error) {
	rta, err = p.pr.GetProductos()

	return rta, err
}

func (p ProductService) GetProductoById(id int) (rta dto.Producto, err error) {
	if id <= 0 {
		return dto.Producto{}, errors.New("el ID no puede ser igual o menor a 0")
	}

	rta, err = p.pr.GetProductoById(id)
	return rta, err
}

func (p ProductService) GetProductsByMinPrice(price float64) (rta []dto.Producto, err error) {
	if price < 0 {
		return nil, errors.New("los precios de los productos no pueden ser negativos")
	}

	rta, err = p.pr.GetProductsByMinPrice(price)
	return rta, err
}

func (p ProductService) SetProducto(newProduct dto.Producto) (savedProd dto.Producto, err error) {
	if err = p.pr.ValidateExistsCodeProduct(newProduct.CodeValue); err != nil {
		return dto.Producto{}, err
	}

	if err = p.validateFormatDate(newProduct.Expiration); err != nil {
		return dto.Producto{}, err
	}

	if newProduct.Price < 0 {
		return dto.Producto{}, errors.New("el precio del producto no puede ser negativo")
	}

	savedProd, err = p.pr.SetProducto(newProduct)
	return savedProd, err
}

func (p ProductService) validateFormatDate(date string) (err error) {
	_, err = time.Parse("02/01/2006", date)
	return err
}
