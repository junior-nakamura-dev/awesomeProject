package main

import (
	"fmt"
	"github.com/google/uuid"
)

type Product struct {
	ID          string
	Description string
	price       float64
}

func NewProduct(description string, price float64) *Product {
	return &Product{
		ID:          uuid.New().String(),
		Description: description,
		price:       price,
	}
}

func main() {
	fmt.Printf("products %+v\n", NewProduct("Test", 10.90))
}
