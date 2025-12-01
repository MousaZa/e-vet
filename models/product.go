package models

type Product struct {
	ID        string  `json:"id"`
	UserId    int     `json:"user_id"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	LastPrice float64 `json:"last_price"`
	IsActive  bool    `json:"is_active"`
}
