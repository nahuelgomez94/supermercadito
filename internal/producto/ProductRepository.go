package producto

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/bootcamp/supermercadito/internal/dto"
)

type productRepository struct {
	productos []dto.Producto
}

func (pr *productRepository) SetProductos() {
	file, err := os.Open("./products.json")

	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(file)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(b, &pr.productos)
}

func (pr productRepository) getProductos() []dto.Producto {
	return pr.productos
}

func (pr productRepository) getProductoById(id int) dto.Producto {
	var rta dto.Producto

	for _, p := range pr.productos {
		if p.Id == id {
			rta = p
			break
		}
	}

	return rta
}

func (pr productRepository) getProductsByMinPrice(price float64) []dto.Producto {
	var rta []dto.Producto

	for _, p := range pr.productos {
		if p.Price >= price {
			rta = append(rta, p)
		}
	}

	return rta
}

func (pr *productRepository) SetProducto(newProduct dto.Producto) (savedProd dto.Producto, err error) {
	savedProd = newProduct

	var id int
	var errCode bool

	for _, p := range pr.productos {
		if p.Id > id {
			id = p.Id
		}

		if p.CodeValue == newProduct.CodeValue {
			errCode = true
			break
		}
	}

	if errCode {
		//c.String(401, "El código del producto ya existe")
		return dto.Producto{}, errors.New("El código del producto ya existe")
	}

	id++
	savedProd.Id = id

	re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")

	if !re.MatchString(savedProd.Expiration) {
		//c.String(401, "Formato de fecha de expiracion incorrecta")
		return dto.Producto{}, errors.New("Formato de fecha de expiracion incorrecta")
	} else {
		reA := strings.Split(savedProd.Expiration, "/")
		dia, errConvDia := strconv.Atoi(reA[0])

		if errConvDia != nil {
			//c.String(401, "No se pudo convertir el día de la fecha de expiracion")
			return dto.Producto{}, errors.New("No se pudo convertir el día de la fecha de expiracion")
		}

		if dia < 0 || dia > 31 {
			//c.String(401, "Dia en fecha de expiracion no valido")
			return dto.Producto{}, errors.New("Dia en fecha de expiracion no valido")
		}
	}

	pr.productos = append(pr.productos, savedProd)

	return savedProd, nil
}
