package models

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	Name string `json:"name"`
	FirstName string `gorm:"column:firstName" json:"firstName"`
	LastName string `gorm:"column:lastName" json:"lastName"`
	Email string `gorm:"column:email" json:"email"`
	Pass string `json:"password"`
	LoggedIn bool `gorm:"column:loggedin" json:"loggedin"`
}

func (Customer) TableName() string {
	return "customers"
}
