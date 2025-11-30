package models

import "math/rand"

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	LastPrice float64 `json:"last_price"`
}

func NewProduct(name string, quantity int, lastPrice float64) *Product {
	return &Product{
		ID:        rand.Intn(100000),
		Name:      name,
		Quantity:  quantity,
		LastPrice: lastPrice,
	}
}
