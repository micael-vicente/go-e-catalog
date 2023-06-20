package models

import (
	"github.com/shopspring/decimal"
)

type ProductsPage struct {
	Items      []Product  `json:"items"`
	Pagination Pagination `json:"pagination"`
}

// Product is a model for a product
type Product struct {
	SKU      string       `json:"sku"`
	Details  Details      `json:"details"`
	Price    Price        `json:"price"`
	Package  Package      `json:"package"`
	Validity Availability `json:"validity"`
}

// Price is a model for product price
type Price struct {
	Amount       decimal.Decimal `json:"amount"`
	PricePerUnit decimal.Decimal `json:"price_per_unit"`
	Currency     string          `json:"currency"`
}

// Details is a model for product details
type Details struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Brand       string `json:"brand"`
	Category    string `json:"category"`
}

// Package is a model for product package
type Package struct {
	Weight float32 `json:"weight"`
	Height float32 `json:"height"`
	Width  float32 `json:"width"`
	Length float32 `json:"length"`
	Type   string  `json:"type"`
	Units  int     `json:"units"`
}

// Availability is a model for product availability
type Availability struct {
	AvailableFrom string `json:"available_from"`
	AvailableTo   string `json:"available_to"`
}

// Pagination is a model for pagination information
type Pagination struct {
	PageNumber    int `json:"page_number"`
	PageSize      int `json:"page_size"`
	TotalElements int `json:"total_elements"`
	TotalPages    int `json:"total_pages"`
}
