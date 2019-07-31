package repositories

import (

	"go-music/models"
)

type IProductRepository interface {
	GetAllProducts() ([]models.Product, error)
	GetPromotionProducts() ([]models.Product, error)
	GetProductById(id int) (models.Product, error)
}

type ProductRepository struct {
	*MusicContext
}

func NewProductRepository(ctx *MusicContext) *ProductRepository {
	return &ProductRepository{
		ctx,
	}
}

func (repo *ProductRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := repo.MusicContext.Find(&products).Error
	return products, err
}

func (repo *ProductRepository) GetPromotionProducts() ([]models.Product, error) {
	var products []models.Product
	err := repo.MusicContext.Where("promotion IS NOT NULL").Find(&products).Error
	return products, err
}

func (repo *ProductRepository) GetProductById(id int) (models.Product, error) {
	var product models.Product
	err := repo.MusicContext.First(&product).Error
	return product, err
}