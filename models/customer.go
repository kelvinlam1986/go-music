package models

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	FirstName string `gorm:"column:first_name"`
	LastName string `gorm:"column:last_name"`
	Email string `gorm:"column:email"`
	Pass string `json:"password"`
	LoggedIn bool `gorm:"column:logged_in"`
}

func (Customer) TableName() string {
	return "customers"
}
