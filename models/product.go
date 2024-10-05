package models

type Product struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	StoreID int     `json:"store_id"`
	StoreName string  `json:"store_name"`

}
