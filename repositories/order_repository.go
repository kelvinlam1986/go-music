package repositories

import "go-music/models"

type IOrderRepository interface {
	GetOrdersByCustomerId(customerId int) ([]models.Order, error)
}
