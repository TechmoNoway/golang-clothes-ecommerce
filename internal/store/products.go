package store

import (
	"context"
	"database/sql"
)

type Product struct {
	ID          int64  `json:"id"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
	Size        string `json:"size"`
	Color       string `json:"color"`
	CategoryID  int64  `json:"category_id"`
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
		product.ProductName,
		product.Description,
		product.Price,
		product.Stock,
		product.Size,
		product.Color,
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

func (s *ProductStore) GetAll(ctx context.Context) ([]Product, error) {

	query := `
		SELECT id, product_name, description, price, stock, size, color, created_at, updated_at 
		FROM products
		WHERE 1 = 1 
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productList []Product

	for rows.Next() {
		var product Product

		err := rows.Scan(
			&product.ID,
			&product.ProductName,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.Size,
			&product.Color,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		productList = append(productList, product)
	}

	return productList, err

}

func (s *ProductStore) GetById(ctx context.Context, productID int64) (*Product, error) {
	query := `
		SELECT id, product_name, description, price, stock, size, color, created_at, updated_at 
		FROM products
		WHERE id = $1 
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	product := &Product{}

	err := s.db.QueryRowContext(
		ctx,
		query,
		productID,
	).Scan(
		&product.ID,
		&product.ProductName,
		&product.Description,
		&product.Price,
		&product.Stock,
		&product.Size,
		&product.Color,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return product, err
}
