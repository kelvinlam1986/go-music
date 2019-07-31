package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Order struct {
	gorm.Model
	Product
	Customer
	CustomerId int `gorm:"column:customer_id"`
	ProductId int `gorm:"column:product_id"`
	Price float64 `gorm:"column:price" json:"sellPrice"`
	PurchaseDate time.Time `gorm:"column:purchase_date" json:"purchaseDate"`
}

func (Order) TableName() string {
	return "orders"
}
