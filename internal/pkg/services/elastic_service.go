package services

import (
	"github.com/elastic/go-elasticsearch/v8"
)

type ElasticService struct {
	Client *elasticsearch.Client
}

func (e *ElasticService) CreateIndexIfNotExists(index string) error {
	_, err := e.Client.Indices.Create(index)
	return err
}

func (e *ElasticService) DeleteIndex(index string) error {
	_, err := e.Client.Indices.Delete([]string{index})
	return err
}
