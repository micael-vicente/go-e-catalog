package db

import (
	"catalog/internal/pkg/models"
)

type ProductRepository interface {
	Get(sku string) (*models.Product, error)
	GetAll(page int, size int) (models.ProductsPage, error)
	Create(p *models.Product) error
	Delete(sku string) error
}
