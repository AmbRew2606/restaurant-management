package models

type MenuItem struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Nutrition   string  `json:"nutrition"`
	Price       float64 `json:"price"`
}
