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

var (
	ErrIdLessZero    = "el ID no puede ser igual o menor a 0"
	ErrPriceLessZero = "los precios de los productos no pueden ser negativos"
	ErrGetProduct    = "Ha ocurrido un problema al obtener el producto a actualizar"
)

func (p ProductService) GetProductos() (rta []dto.Producto, err error) {
	rta, err = p.pr.GetProductos()

	return rta, err
}

func (p ProductService) GetProductoById(id int) (rta dto.Producto, err error) {
	if id <= 0 {
		return dto.Producto{}, errors.New(ErrIdLessZero)
	}

	rta, err = p.pr.GetProductoById(id)
	return rta, err
}

func (p ProductService) GetProductsByMinPrice(price float64) (rta []dto.Producto, err error) {
	if err := p.validatePrice(price); err != nil {
		return nil, err
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

	if err = p.validatePrice(newProduct.Price); err != nil {
		return dto.Producto{}, err
	}

	savedProd, err = p.pr.SetProducto(newProduct)
	return savedProd, err
}

func (p ProductService) UpdateProduct(id int, prod dto.ProductoRequest) (rta dto.Producto, err error) {
	_, err = p.GetProductoById(id)

	if err != nil {
		return dto.Producto{}, err
	}

	err = p.validateUniqueCode(id, prod.CodeValue)

	if err != nil {
		return dto.Producto{}, err
	}

	var prodDTO dto.Producto
	prodDTO.SetValues(prod)
	prodDTO.Id = id

	savedProd, err := p.pr.UpdateProduct(id, prodDTO)

	if err != nil {
		return dto.Producto{}, err
	}

	return savedProd, nil
}

func (p ProductService) validateFormatDate(date string) (err error) {
	_, err = time.Parse("02/01/2006", date)
	return err
}

func (p ProductService) validateUniqueCode(id int, code string) (err error) {
	return p.pr.ValidateUniqueCode(id, code)
}

func (p ProductService) PatchProduct(id int, producto dto.Producto) (savedProdu dto.Producto, err error) {
	prodNoUpdated, err := p.pr.GetProductoById(id)

	if err != nil {
		return dto.Producto{}, errors.New(ErrGetProduct)
	}

	if producto.Expiration != "" {
		err := p.validateFormatDate(producto.Expiration)
		if err != nil {
			return dto.Producto{}, err
		}
		prodNoUpdated.Expiration = producto.Expiration
	}

	if producto.CodeValue != "" {
		prodNoUpdated.CodeValue = producto.CodeValue
	}

	if producto.IsPublished != nil {
		prodNoUpdated.IsPublished = producto.IsPublished
	}

	if producto.Name != "" {
		prodNoUpdated.Name = producto.Name
	}

	if producto.Price != 0 {
		if err := p.validatePrice(producto.Price); err != nil {
			return dto.Producto{}, err
		}
		prodNoUpdated.Price = producto.Price
	}

	if producto.Quantity != 0 {
		prodNoUpdated.Quantity = producto.Quantity
	}

	savedProduct, err := p.pr.UpdateProduct(id, prodNoUpdated)

	if err != nil {
		return dto.Producto{}, err
	}

	return savedProduct, nil
}

func (p ProductService) validatePrice(price float64) (err error) {
	if price < 0 {
		return errors.New(ErrPriceLessZero)
	}

	return nil
}
