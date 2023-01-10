package dto

type ProductoRequest struct {
	Id          int     `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	CodeValue   string  `json:"code_value" binding:"required"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}

type Producto struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished *bool   `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (p *Producto) SetValues(prod ProductoRequest) (prodResult Producto, err error) {
	p.Name = prod.Name
	p.Quantity = prod.Quantity
	p.CodeValue = prod.CodeValue
	p.IsPublished = &prod.IsPublished
	p.Expiration = prod.Expiration
	p.Price = prod.Price
	return *p, nil
}
