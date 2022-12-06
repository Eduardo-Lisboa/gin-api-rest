package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(100);not null;unique" validate:"nonzero"`
	Password string `json:"password" gorm:"type:varchar(100);not null" validate:"nonzero"`
}

var Users []User
