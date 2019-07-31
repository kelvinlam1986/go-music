package repositories

import "go-music/models"

type ICustomerRepository interface {
	GetCustomerById(id int) (models.Customer, error)
	GetCustomerByName(firstName string, lastName string) (models.Customer, error)
	AddUser(customer models.Customer) (models.Customer, error)
	SignIn(username, password string) (models.Customer, error)
	SignOut(userId int) error
}
