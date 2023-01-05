package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/bootcamp/supermercadito/internal/dto"
	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id          int     `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	CodeValue   string  `json:"code_value" binding:"required"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}

func SetProductGroupRoutes(server *gin.Engine) {
	p := server.Group("/products")

	p.GET("/", func(c *gin.Context) {
		c.JSON(200, GetProductos())
	})

	p.GET("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.String(500, "Error en la conversion ID")
			return
		}
		if id <= 0 {
			c.String(500, "No se ha especificado ID")
			return
		}

		rta := GetProductoById(id)

		if rta.Id != id {
			c.String(500, "No se ha encontrado el ID especificado")
			return
		}

		c.JSON(200, rta)
	})

	p.GET("/search", func(c *gin.Context) {
		pPrice := c.Query("pricegt")

		fmt.Println(pPrice)
		price, err := strconv.ParseFloat(pPrice, 10)

		if err != nil {
			c.String(500, err.Error())
			return
		}

		if price < 0 {
			c.String(500, "El precio no puede ser negativo")
			return
		}

		rta := GetProductsByMinPrice(price)

		if len(rta) == 0 {
			c.String(200, "No hay productos con ese precio")
			return
		}

		c.JSON(200, rta)

	})

	p.POST("/", func(c *gin.Context) {
		var prodReq Producto
		err := c.ShouldBindJSON(&prodReq)

		if err != nil {
			c.String(401, err.Error())
			return
		}

		var id int
		var errCode bool
		for _, p := range productos {
			if p.Id > id {
				id = p.Id
			}

			if p.CodeValue == prodReq.CodeValue {
				errCode = true
				break
			}
		}

		if errCode {
			c.String(401, "El código del producto ya existe")
			return
		}

		id++
		prodReq.Id = id

		re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")

		if !re.MatchString(prodReq.Expiration) {
			c.String(401, "Formato de fecha de expiracion incorrecta")
			return
		} else {
			reA := strings.Split(prodReq.Expiration, "/")
			dia, errConvDia := strconv.Atoi(reA[0])

			if errConvDia != nil {
				c.String(401, "No se pudo convertir el día de la fecha de expiracion")
				return
			}

			if dia < 0 || dia > 31 {
				c.String(401, "Dia en fecha de expiracion no valido")
				return
			}
		}

		c.JSON(200, prodReq)
	})
}

var productos []dto.Producto

func SetProductos() {
	file, err := os.Open("./products.json")

	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(file)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(b, &productos)

	fmt.Println(productos)
}

func GetProductos() []dto.Producto {
	return productos
}

func GetProductoById(id int) dto.Producto {
	var rta dto.Producto

	for _, p := range productos {
		if p.Id == id {
			rta = p
			break
		}
	}

	return rta
}

func GetProductsByMinPrice(price float64) []dto.Producto {
	var rta []dto.Producto

	for _, p := range productos {
		if p.Price >= price {
			rta = append(rta, p)
		}
	}

	return rta
}
