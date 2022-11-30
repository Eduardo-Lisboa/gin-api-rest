package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string  `json:"code" validate:"nonzero"`
	Name  string  `json:"name" validate:"nonzero"`
	Price float64 `json:"price" validate:"nonzero"`
}

var Products []Product

func Validate(porduct *Product) error {
	if err := validator.Validate(porduct); err != nil {
		return err
	}
	return nil

}
