package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Code  string `json:"code"`
	Price uint   `json:"price"`
}
