package models

type Store struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Location string `json:"location"`
}
