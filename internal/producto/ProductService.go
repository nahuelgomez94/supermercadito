package producto

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/bootcamp/supermercadito/internal/dto"
)

var pr productRepository

func SetProductos() {
	pr.setProductos()
}

func GetProductos() (rta []dto.Producto, err error) {
	rta, err = pr.getProductos()
	return rta, err
}

func GetProductoById(id int) (rta dto.Producto, err error) {
	if id <= 0 {
		return dto.Producto{}, errors.New("El ID no puede ser igual o menor a 0")
	}

	rta, err = pr.getProductoById(id)
	return rta, err
}

func GetProductsByMinPrice(price float64) (rta []dto.Producto, err error) {
	if price < 0 {
		return nil, errors.New("Los precios de los productos no pueden ser negativos")
	}

	rta, err = pr.getProductsByMinPrice(price)
	return rta, err
}

func SetProducto(newProduct dto.Producto) (savedProd dto.Producto, err error) {
	if err = pr.validateExistsCodeProduct(newProduct.CodeValue); err != nil {
		return dto.Producto{}, err
	}

	if err = validateFormatDate(newProduct.Expiration); err != nil {
		return dto.Producto{}, err
	}

	if newProduct.Price < 0 {
		return dto.Producto{}, errors.New("El precio del producto no puede ser negativo")
	}

	savedProd, err = pr.setProducto(newProduct)
	return savedProd, err
}

func validateFormatDate(date string) (err error) {
	//re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
	re := regexp.MustCompile("([0-9]{2})([/])([0-9]{2})([/])([0-9]{4})")
	if !re.MatchString(date) {
		//c.String(401, "Formato de fecha de expiracion incorrecta")
		return errors.New("Formato de fecha de expiracion incorrecta")
	} else {
		reA := strings.Split(date, "/")
		dia, errConvDia := strconv.Atoi(reA[0])
		mes, errConvMes := strconv.Atoi(reA[1])

		if errConvDia != nil || errConvMes != nil {
			//c.String(401, "No se pudo convertir el día de la fecha de expiracion")
			return errors.New("No se pudo convertir el día de la fecha de expiracion")
		}

		if dia < 0 || dia > 31 {
			//c.String(401, "Dia en fecha de expiracion no valido")
			return errors.New("Dia en fecha de expiracion no valido")
		}

		if mes > 12 || mes < 0 {
			return errors.New("Mes en fecha de expiracion no valido")
		}
	}

	return nil
}
