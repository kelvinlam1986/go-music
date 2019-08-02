package repositories

import "go-music/models"

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