package controllers

import (
	"github.com/gin-gonic/gin"
	"go-music/models"
	"go-music/repositories"
	"net/http"
)

type CustomerController struct {
	CustomerRepository repositories.ICustomerRepository
}

func NewCustomerController(customerRepository repositories.ICustomerRepository) *CustomerController {
	return &CustomerController{ CustomerRepository: customerRepository }
}

func (controller *CustomerController) SignIn(context *gin.Context) {
	var customer models.Customer
	err := context.ShouldBindJSON(&customer)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "error": err })
		return
	}

	customer, err = controller.CustomerRepository.SignIn(customer.Email, customer.Pass)
	if err != nil {
		if err == repositories.ErrorInvalidPassword {
			context.JSON(http.StatusForbidden, gin.H{ "error": err })
			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{ "error": err })
		return
	}

	context.JSON(http.StatusOK, customer)
}
