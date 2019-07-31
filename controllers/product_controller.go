package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-music/repositories"
	"net/http"
)

type ProductController struct {
	ProductRepository repositories.IProductRepository
}

func NewProductController(productRepository repositories.IProductRepository) *ProductController {
	return &ProductController{ ProductRepository:  productRepository}
}

func (controller *ProductController) GetAllProducts(context * gin.Context) {
	products, err := controller.ProductRepository.GetAllProducts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Found %d products\n", len(products))
	context.JSON(http.StatusOK, products)
}
