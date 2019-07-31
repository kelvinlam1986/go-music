package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Image string `json:"img"`
	ImageAlt string `gorm:"column:imgalt" json:"imgAlt"`
	Price float64 `json:"price"`
	Promotion float64 `json:"promotion"`
	ProductName string `gorm:"column:productname" json:"productName"`
	Description string `json:"description"`
}

func (Product) TableName() string {
	return "products"
}
