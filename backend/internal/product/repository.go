// Root: clothing-store/backend/internal/product/repository.go
package product

import (
	"context"
	"database/sql"
)

func GetAllProducts(ctx context.Context, db *sql.DB) ([]Product, error) {
	rows, err := db.QueryContext(ctx, `
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
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Category, &p.Price, &p.ImageURL, &p.InStock, &p.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func GetProductByID(ctx context.Context, db *sql.DB, id int) (*Product, error) {
	row := db.QueryRowContext(ctx, `
        SELECT id, name, description, category, price, image_url, in_stock, created_at
        FROM products
        WHERE id = $1
    `, id)

	var p Product
	if err := row.Scan(&p.ID, &p.Name, &p.Description, &p.Category, &p.Price, &p.ImageURL, &p.InStock, &p.CreatedAt); err != nil {
		return nil, err
	}
	return &p, nil
}
