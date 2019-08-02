package repositories

import (
	"go-music/models"
	"go-music/utils"
)

type ICustomerRepository interface {
	GetCustomerById(id int) (models.Customer, error)
	GetCustomerByName(firstName string, lastName string) (models.Customer, error)
	AddUser(customer models.Customer) (models.Customer, error)
	SignIn(username, password string) (models.Customer, error)
	SignOut(userId int) error
}

type CustomerRepository struct {
	*MusicContext
}

func NewCustomerRepository(ctx *MusicContext) *CustomerRepository {
	return &CustomerRepository{ MusicContext: ctx }
}

func (repo *CustomerRepository) GetCustomerById(id int) (models.Customer, error) {
	var customer models.Customer
	err := repo.MusicContext.First(&customer, id).Error
	return customer, err
}

func (repo *CustomerRepository) GetCustomerByName(firstName string, lastName string) (models.Customer, error) {
	var customer models.Customer
	err := repo.MusicContext.Where(&models.Customer{ FirstName: firstName, LastName: lastName }).
			Find(&customer).Error
	return customer, err
}

func (repo *CustomerRepository) AddUser(customer models.Customer) (models.Customer, error) {
	utils.HashPassword(&customer.Pass)
	customer.LoggedIn = true
	customer.Pass = ""
	return customer, repo.MusicContext.Create(&customer).Error
}


func (repo *CustomerRepository) SignIn(username, password string) (models.Customer, error) {
	var customer models.Customer
	// Obtain a *gorm.DB object representing our customer's row
	result := repo.MusicContext.Table("customers").Where(&models.Customer{ Email: username })
	// Retrieve the data for the customer with the passed email
	err := result.First(&customer).Error
	if err != nil {
		return customer, err
	}

	if  !utils.CheckPassword(customer.Pass, password) {
		return customer, ErrorInvalidPassword
	}

	customer.Pass = ""
	err = result.Update("loggedin", 1).Error
	if err != nil {
		return customer, err
	}

	return customer, result.Find(&customer).Error
}

func (repo *CustomerRepository) SignOut(userId int) error {
	return nil
}