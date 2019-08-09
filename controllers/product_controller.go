package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-music/repositories"
	"go-music/viewmodels"
	"net/http"
)

type ProductController struct {
	ProductRepository repositories.IProductRepository
}

func NewProductController(productRepository repositories.IProductRepository) *ProductController {
	return &ProductController{ ProductRepository:  productRepository}
}

func (controller *ProductController) GetAllProducts(context * gin.Context) {
	var productsVm []viewmodels.ProductGetAllVm
	products, err := controller.ProductRepository.GetAllProducts()
	for _, product := range products {
		productsVm = append(productsVm, viewmodels.ProductGetAllVm{
			Id: product.ID,
			Description: product.Description,
			ProductName: product.ProductName,
			Image: product.Image,
			ImageAlt: product.ImageAlt,
			Price: product.Price,
			Promotion: product.Promotion,
		});
	}
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Found %d products\n", len(products))
	context.JSON(http.StatusOK, productsVm)
}
