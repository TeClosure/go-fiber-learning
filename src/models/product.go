package models

type Product struct {
	ID			uint `json:"id"`
	Title		string `json:"title"`
	Description	string `json:"description"`
	Image		string `json:"image"`
	Price		string `json:"price"`
}