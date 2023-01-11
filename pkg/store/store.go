package store

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"github.com/bootcamp/supermercadito/internal/dto"
	"github.com/bootcamp/supermercadito/internal/interfaces"
)

type ProductStore struct {
	interfaces.IProductStorage
}

func NewProductStorage() *ProductStore {
	return &ProductStore{}
}

func (p *ProductStore) Get() (productos []dto.Producto, err error) {
	file, err := os.Open("./products.json")
	//file, err := os.Open("../products.json")

	if err != nil {
		return nil, err
	}

	defer file.Close()

	b, err := io.ReadAll(file)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &productos)

	/*for _, p := range productos {
		if p.Id > pr.idMax {
			pr.idMax = p.Id
		}
	}*/

	return productos, err
}

func (p *ProductStore) Set(productos []dto.Producto) (err error) {
	data, err := json.Marshal(productos)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./products.json", data, 0)

	return err
}
