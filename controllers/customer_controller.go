package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-music/models"
	"go-music/repositories"
	"go-music/viewmodels"
	"net/http"
	"strconv"
)

type CustomerController struct {
	CustomerRepository repositories.ICustomerRepository
}

func NewCustomerController(customerRepository repositories.ICustomerRepository) *CustomerController {
	return &CustomerController{ CustomerRepository: customerRepository }
}

func (controller *CustomerController) GetCustomerById(context *gin.Context) {
	idParam := context.Param("id")
	customerId, err := strconv.Atoi(idParam)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	customer, err := controller.CustomerRepository.GetCustomerById(customerId);
	if err != nil {
		msg := fmt.Sprintf("Customer with id %s not found", idParam)
		context.JSON(http.StatusNotFound, gin.H{"error": msg})
		return
	}

	customerVm := viewmodels.CustomerGetByIdVm{
		Id: customer.ID,
		FirstName: customer.FirstName,
		LastName: customer.LastName,
		Email: customer.Email,
		LoggedIn: customer.LoggedIn,
	}

	context.JSON(http.StatusOK, customerVm)
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


