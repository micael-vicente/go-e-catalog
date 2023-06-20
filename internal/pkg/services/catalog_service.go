package services

import (
	"catalog/internal/pkg/db"
	"catalog/internal/pkg/models"
)

type CatalogService struct {
	Repo db.ProductRepository
}

// GetProduct returns a product by sku
func (cs *CatalogService) GetProduct(sku string) (*models.Product, error) {
	return cs.Repo.Get(sku)
}

// GetAllProducts returns all products
func (cs *CatalogService) GetAllProducts(page int, size int) (models.ProductsPage, error) {
	return cs.Repo.GetAll(page, size)
}

// CreateProduct creates a new product
func (cs *CatalogService) CreateProduct(p *models.Product) error {
	return cs.Repo.Create(p)
}

// DeleteProduct deletes a product
func (cs *CatalogService) DeleteProduct(sku string) error {
	return cs.Repo.Delete(sku)
}
