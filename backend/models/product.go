package models

import (
	"clothing-store/backend/utils"
	"context"
	"time"
)

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Price       float64   `json:"price"`
	ImageURL    string    `json:"image_url"`
	InStock     bool      `json:"in_stock"`
	CreatedAt   time.Time `json:"created_at"`
}

func GetAllProducts() ([]Product, error) {
	rows, err := utils.DB.Query(context.Background(), `
        SELECT id, name, description, category, price, image_url, in_stock, created_at
        FROM products
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Category, &p.Price, &p.ImageURL, &p.InStock, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func GetProductByID(id int) (*Product, error) {
	row := utils.DB.QueryRow(context.Background(), `
        SELECT id, name, description, category, price, image_url, in_stock, created_at
        FROM products
        WHERE id = $1
    `, id)

	var p Product
	err := row.Scan(&p.ID, &p.Name, &p.Description, &p.Category, &p.Price, &p.ImageURL, &p.InStock, &p.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
