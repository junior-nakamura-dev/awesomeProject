package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
	db, err := sql.Open("mysql", "root:example@tcp(localhost:3306)/go_database")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	product := NewProduct("Test I", 10.90)
	err = insertProduct(db, *product)

	if err != nil {
		panic(err)
	}

	product1 := NewProduct("Test I", 10.90)
	err = insertProduct(db, *product1)

	if err != nil {
		panic(err)
	}

	product.Description = "Updated product I"
	updateProduct(db, *product)

	product, err = selectProduct(db, product.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println(">" + product.Description + "<")

	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}

	for _, product := range products {
		fmt.Println(">" + product.Description + "<")
	}

	defer db.Close()
}

func insertProduct(db *sql.DB, product Product) error {

	stm, err := db.Prepare("INSERT INTO PRODUCTS (id, description, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(product.ID, product.Description, product.price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product Product) error {
	stm, err := db.Prepare("UPDATE PRODUCTS SET description = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(product.Description, product.price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stm, err := db.Prepare("SELECT id, description, price FROM PRODUCTS WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer stm.Close()
	var product Product
	err = stm.QueryRow(id).Scan(&product.ID, &product.Description, &product.price)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
func selectAllProducts(db *sql.DB) ([]Product, error) {
	stm, err := db.Prepare("SELECT id, description, price FROM PRODUCTS")
	if err != nil {
		return nil, err
	}
	defer stm.Close()

	rows, err := stm.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Description, &product.price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
