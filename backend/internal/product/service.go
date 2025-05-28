package product

import (
	"context"
	"database/sql"
)

func ListProducts(ctx context.Context, db *sql.DB) ([]Product, error) {
	return GetAllProducts(ctx, db)
}

func GetProduct(ctx context.Context, db *sql.DB, id int) (*Product, error) {
	return GetProductByID(ctx, db, id)
}
