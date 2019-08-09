package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Image string
	ImageAlt string `gorm:"column:image_alt"`
	Price float64 `json:"price"`
	Promotion float64 `json:"promotion"`
	ProductName string `gorm:"column:product_name"`
	Description string `json:"description"`
}

func (Product) TableName() string {
	return "products"
}
