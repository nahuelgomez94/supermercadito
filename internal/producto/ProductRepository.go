package producto

import (
	"errors"
	"strconv"

	"github.com/bootcamp/supermercadito/internal/dto"
	"github.com/bootcamp/supermercadito/internal/interfaces"
)

type ProductRepository struct {
	interfaces.IProductRepository
	Storage interfaces.IProductStorage
}

func NewProductRepository(storage interfaces.IProductStorage) *ProductRepository {
	rta := &ProductRepository{}
	rta.Storage = storage
	return rta
}

func (pr *ProductRepository) getNewId() (id int) {
	productos, _ := pr.Storage.Get()

	id = 0
	for _, p := range productos {
		if id < p.Id {
			id = p.Id
		}
	}

	id++
	return id
}

func (pr ProductRepository) GetProductos() (rta []dto.Producto, err error) {
	productos, err := pr.Storage.Get()

	return productos, err
}

func (pr ProductRepository) GetProductoById(id int) (rta dto.Producto, err error) {
	productos, err := pr.Storage.Get()

	if err != nil {
		return dto.Producto{}, err
	}

	for _, p := range productos {
		if p.Id == id {
			rta = p
			break
		}
	}

	if rta.Id != id {
		return rta, errors.New("No se ha encontrado el recurso")
	}

	return rta, nil
}

func (pr ProductRepository) GetProductsByMinPrice(price float64) (rta []dto.Producto, err error) {
	productos, err := pr.Storage.Get()

	if err != nil {
		return nil, err
	}

	for _, p := range productos {
		if p.Price >= price {
			rta = append(rta, p)
		}
	}

	return rta, nil
}

func (pr *ProductRepository) ValidateExistsCodeProduct(code string) (err error) {
	productos, err := pr.Storage.Get()

	if err != nil {
		return err
	}

	for _, p := range productos {
		if p.CodeValue == code {
			err = errors.New("el còdigo ya existe en la base de datos")
			break
		}
	}

	return err
}

func (pr *ProductRepository) SetProducto(newProduct dto.Producto) (savedProd dto.Producto, err error) {
	savedProd = newProduct
	savedProd.Id = pr.getNewId()

	productos, err := pr.Storage.Get()

	if err != nil {
		return dto.Producto{}, err
	}

	productos = append(productos, savedProd)

	err = pr.Storage.Set(productos)
	return savedProd, err
}

func (pr *ProductRepository) UpdateProduct(id int, prod dto.Producto) (savedProd dto.Producto, err error) {
	productos, err := pr.Storage.Get()

	for i, p := range productos {
		if p.Id == id {
			productos[i] = prod

			err = pr.Storage.Set(productos)

			return productos[i], err
		}
	}

	return dto.Producto{}, errors.New("No existe el producto en la base de datos")
}

func (pr *ProductRepository) ValidateUniqueCode(id int, code string) (err error) {
	productos, err := pr.Storage.Get()

	if err != nil {
		return err
	}

	for _, p := range productos {
		if id != p.Id && code == p.CodeValue {
			msg := "El código ya existe en la base de datos. Producto ID: " + strconv.FormatInt(int64(p.Id), 10)
			return errors.New(msg)
		}
	}

	return nil
}

func (pr *ProductRepository) DeleteProduct(id int) (err error) {
	productos, err := pr.Storage.Get()

	if err != nil {
		return err
	}

	for i, p := range productos {
		if p.Id == id {
			productos = append(productos[:i], productos[i+1:]...)

			err = pr.Storage.Set(productos)
			return err
		}
	}

	return errors.New("No se encontró el producto indicado")
}
