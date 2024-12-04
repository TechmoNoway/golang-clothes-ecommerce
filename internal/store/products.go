package store

import (
	"context"
	"database/sql"
)

type Product struct {
	ID          int64  `json:"id"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Size        string `json:"size"`
	Color       string `json:"color"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ProductStore struct {
	db *sql.DB
}

func (s *ProductStore) create(ctx context.Context, product *Product) error {
	query := `INSERT INTO products (product_name, description, price, stock, size, color) VALUES ($1, $2, $3, $4, $5, $6)`

	err := s.db.QueryRowContext(
		ctx,
		query,
		&product.ProductName,
		&product.Description,
		&product.Price,
		&product.Stock,
		&product.Size,
		&product.Color,
	).Scan(
		&product.ID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
