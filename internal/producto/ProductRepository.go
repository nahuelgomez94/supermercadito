package producto

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/bootcamp/supermercadito/internal/dto"
	"github.com/bootcamp/supermercadito/internal/interfaces"
)

type ProductRepository struct {
	interfaces.IProductRepository
	productos []dto.Producto
	idMax     int
}

func NewProductRepository() *ProductRepository {
	rta := &ProductRepository{}
	rta.SetProductos()
	return rta
}

func (pr *ProductRepository) getNewId() (id int) {
	pr.idMax++
	id = pr.idMax
	return id
}

func (pr *ProductRepository) SetProductos() (err error) {
	file, err := os.Open("./products.json")

	if err != nil {
		return err
	}

	b, err := io.ReadAll(file)

	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &pr.productos)

	for _, p := range pr.productos {
		if p.Id > pr.idMax {
			pr.idMax = p.Id
		}
	}

	return err
}

func (pr ProductRepository) GetProductos() (rta []dto.Producto, err error) {
	return pr.productos, nil
}

func (pr ProductRepository) GetProductoById(id int) (rta dto.Producto, err error) {
	for _, p := range pr.productos {
		if p.Id == id {
			rta = p
			break
		}
	}

	return rta, nil
}

func (pr ProductRepository) GetProductsByMinPrice(price float64) (rta []dto.Producto, err error) {
	for _, p := range pr.productos {
		if p.Price >= price {
			rta = append(rta, p)
		}
	}

	return rta, nil
}

func (pr *ProductRepository) ValidateExistsCodeProduct(code string) (err error) {
	for _, p := range pr.productos {

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
	pr.productos = append(pr.productos, savedProd)

	return savedProd, nil
}
