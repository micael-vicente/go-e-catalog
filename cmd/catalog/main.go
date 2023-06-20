package main

import (
	"catalog/config"
	"catalog/internal/pkg/controllers"
	"catalog/internal/pkg/db"
	"catalog/internal/pkg/services"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

func main() {

	conf := getConfig("dev")

	engine := gin.Default()
	elastic := getElastic(conf.Elastic)

	initializeES(elastic, conf.Elastic.Index)

	repo := &db.ProductRepositoryImpl{Client: elastic, ProductsIndex: conf.Elastic.Index}
	service := services.CatalogService{Repo: repo}
	controller := controllers.CatalogController{Service: service}

	engine.GET("/products", controller.ListProducts)
	engine.GET("/products/:sku", controller.GetProduct)
	engine.POST("/products", controller.CreateProduct)
	engine.DELETE("/products/:sku", controller.DeleteProduct)

	err := engine.Run(":" + strconv.Itoa(conf.Server.Port))

	if err != nil {
		log.Fatal(err)
	}
}

func getConfig(env string) config.Config {
	viper.SetConfigName("config-" + env)
	viper.AddConfigPath("config")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	var conf config.Config

	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatal(err)
	}
	return conf
}

// getElastic returns an elasticsearch client
func getElastic(config config.ElasticConfig) *elasticsearch.Client {
	esConfig := elasticsearch.Config{
		Addresses: []string{
			config.Url,
		},
		Username: config.Username,
		Password: config.Password,
	}

	if client, err := elasticsearch.NewClient(esConfig); err != nil {
		log.Fatal(err)
		return nil
	} else {
		return client
	}
}

// initializeES creates the index if it doesn't exist
func initializeES(client *elasticsearch.Client, index string) {
	service := services.ElasticService{Client: client}

	err := service.CreateIndexIfNotExists(index)
	if err != nil {
		return
	}
}
