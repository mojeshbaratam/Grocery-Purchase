package models

type Rating struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	ProductID   int    `json:"product_id"`
	Rating      int    `json:"rating"` // Between 1 and 5
	Description string `json:"description"`
}