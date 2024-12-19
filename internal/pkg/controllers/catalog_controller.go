package controllers

import (
	"catalog/internal/pkg/models"
	"catalog/internal/pkg/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CatalogController struct {
	Service services.CatalogService
}

// ListProducts lists all products
func (cc *CatalogController) ListProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	if size <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "size must be greater than 0"})
		return
	}

	products, err := cc.Service.GetAllProducts(page, size)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, products)
	}
}

// GetProduct gets a product by sku
func (cc *CatalogController) GetProduct(c *gin.Context) {
	product, err := cc.Service.GetProduct(c.Param("sku"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, product)
	}
}

// CreateProduct creates a new product (or replaces existing)
func (cc *CatalogController) CreateProduct(c *gin.Context) {
	var product models.Product
	err := c.ShouldBindJSON(&product)
	err = cc.Service.CreateProduct(&product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, product)
	}
}

// DeleteProduct deletes a product by sku
func (cc *CatalogController) DeleteProduct(c *gin.Context) {
	err := cc.Service.DeleteProduct(c.Param("sku"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusNoContent, "")
	}
}
