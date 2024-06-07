package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string `gorm:"column:name"`
	Price int    `gorm:"column:price"`
}

func (product *Product) TableName() string {
	return "product"
}
