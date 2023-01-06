package producto

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/bootcamp/supermercadito/internal/dto"
)

type productRepository struct {
	productos []dto.Producto
}

var idMax int

func getNewId() (id int) {
	idMax++
	id = idMax
	return id
}

func (pr *productRepository) setProductos() (err error) {
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
		if p.Id > idMax {
			idMax = p.Id
		}
	}

	return err
}

func (pr productRepository) getProductos() (rta []dto.Producto, err error) {
	return pr.productos, nil
}

func (pr productRepository) getProductoById(id int) (rta dto.Producto, err error) {
	for _, p := range pr.productos {
		if p.Id == id {
			rta = p
			break
		}
	}

	return rta, nil
}

func (pr productRepository) getProductsByMinPrice(price float64) (rta []dto.Producto, err error) {
	for _, p := range pr.productos {
		if p.Price >= price {
			rta = append(rta, p)
		}
	}

	return rta, nil
}

func (pr *productRepository) validateExistsCodeProduct(code string) (err error) {
	for _, p := range pr.productos {

		if p.CodeValue == code {
			err = errors.New("el còdigo ya existe en la base de datos")
			break
		}
	}

	return err
}

func (pr *productRepository) setProducto(newProduct dto.Producto) (savedProd dto.Producto, err error) {
	savedProd = newProduct
	savedProd.Id = getNewId()
	pr.productos = append(pr.productos, savedProd)

	return savedProd, nil
}
