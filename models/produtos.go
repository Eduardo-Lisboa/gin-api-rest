package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var Products []Product
