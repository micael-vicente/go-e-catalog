# Go e-catalog

A very simple cart service written in Go.
The service is a REST API that allows you to create, list, update and delete carts and items.

## Requirements 📖
- Go 1.23.4 
- Docker
- Docker Compose

## Running ▶️

### Start the service

```bash
docker-compose -f docker/docker-compose.yml up
```

```bash
go run cmd/catalog/main.go
```

### Stopping the service

Send a `SIGINT` signal to the process.

```bash
docker-compose -f docker/docker-compose.yml down
```

## Usage

### Create/Update a product

```bash
curl -X POST http://localhost:8081/products \
-d '{"sku": "SKU1"}' \
-H "Content-Type: application/json"
```

See below example for the whole schema
<details>
    <summary>Product model example</summary>
    
    {
        "sku": "SKU11",
        "details": {
            "name": "MyJUICE 330ml",
            "description": "MyJUICE 330ml Cans Carton",
            "brand": "MyJUICE",
            "category": "Carbonated Soft Drink"
        },
        "price": {
            "value": "5",
            "per_unit": "0.5",
            "currency": "EUR"
        },
        "package": {
            "weight": "500 g",
            "height": "10 cm",
            "width": "25 cm",
            "length": "10 cm",
            "type": "Carton",
            "units": 10
        },
        "validity": {
            "available_from": "1930-01-05",
            "available_to": "3000-01-05"
        }
    }
</details>

### List all products

If not provided, `page` and `size` will default to `0` and `10`, respectively.  

```bash
curl -X GET http://localhost:8081/products?size=10&page=0
```

### Get a product by SKU

```bash
curl -X GET http://localhost:8081/products/SKU1
```

### Delete a product by SKU

```bash
curl -X DELETE http://localhost:8081/products/SKU1
```