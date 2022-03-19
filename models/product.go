package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID string `json: "id"`
	Name string `json: "name"`
	Brand string `json: "brand"`
	Store string `json "store"`
	Price float64 `json: "price"`
	Quantity int `json: quantity`
}