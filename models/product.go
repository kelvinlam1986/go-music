package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Image string
	ImageAlt string `gorm:"column:image_alt"`
	Price float64
	Promotion float64
	ProductName string `gorm:"column:product_name"`
	Description string
}

func (Product) TableName() string {
	return "products"
}
