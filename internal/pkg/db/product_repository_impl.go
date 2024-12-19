package db

import (
	"catalog/internal/pkg/models"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

type ProductRepositoryImpl struct {
	Client        *elasticsearch.Client
	ProductsIndex string
}

// Get returns a product by sku
func (r *ProductRepositoryImpl) Get(sku string) (*models.Product, error) {
	if doc, err := r.Client.Get(r.ProductsIndex, sku); err != nil {
		return nil, err
	} else {
		var body models.ElasticSource
		err := json.NewDecoder(doc.Body).Decode(&body)
		return &body.Source, err
	}
}

func (r *ProductRepositoryImpl) GetAll(page int, size int) (models.ProductsPage, error) {
	search, err := r.Client.Search(
		r.Client.Search.WithIndex(r.ProductsIndex),
		r.Client.Search.WithSize(size),
		r.Client.Search.WithFrom(getFrom(page, size)),
	)

	var response models.ElasticPagedResponse
	err = json.NewDecoder(search.Body).Decode(&response)

	var productsPage models.ProductsPage

	//items
	productsPage.Items = []models.Product{}
	for _, product := range response.Hits.Docs {
		productsPage.Items = append(productsPage.Items, product.Source)
	}

	//pagination
	productsPage.Pagination.PageNumber = page
	productsPage.Pagination.TotalElements = response.Hits.Total.Value
	productsPage.Pagination.TotalPages = getTotalPages(size, response.Hits.Total.Value)
	productsPage.Pagination.PageSize = len(productsPage.Items)

	return productsPage, err
}

func (r *ProductRepositoryImpl) Create(p *models.Product) error {
	_, err := r.Client.Index(
		r.ProductsIndex,
		esutil.NewJSONReader(&p),
		r.Client.Index.WithDocumentID(p.SKU),
	)
	return err
}

func (r *ProductRepositoryImpl) Delete(id string) error {
	_, err := r.Client.Delete(r.ProductsIndex, id)
	return err
}

func getFrom(page int, size int) int {
	if page <= 0 {
		return 0
	} else {
		return page * size
	}
}

func getTotalPages(size int, totalElements int) int {
	if totalElements == 0 {
		return 0
	}

	pages := totalElements / size

	if extra := totalElements%size > 0; extra {
		return pages + 1
	} else {
		return pages
	}
}
