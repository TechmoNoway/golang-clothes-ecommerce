package store

import (
	"context"
	"database/sql"
	"errors"
)

type Product struct {
	ID          int64    `json:"id"`
	ProductName string   `json:"product_name"`
	Description string   `json:"description"`
	Price       int64    `json:"price"`
	Stock       int64    `json:"stock"`
	Size        string   `json:"size"`
	Color       string   `json:"color"`
	CategoryID  int64    `json:"category_id"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	Category    Category `json:"category"`
}

type ProductStore struct {
	db *sql.DB
}

func (s *ProductStore) Create(ctx context.Context, product *Product) error {
	query := `
	INSERT INTO products (product_name, description, price, stock, size, color, category_id) VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id, created_at, updated_at
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		product.ProductName,
		product.Description,
		product.Price,
		product.Stock,
		product.Size,
		product.Color,
		product.CategoryID,
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
		SELECT products.id, product_name, description, price, stock, size, color, products.created_at, products.updated_at, categories.*
		FROM products
		JOIN categories ON (products.category_id = categories.id)
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
			&product.Category.ID,
			&product.Category.CategoryName,
			&product.Category.CreatedAt,
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
		SELECT products.id, product_name, description, price, stock, size, color, products.created_at, products.updated_at, categories.*
		FROM products
		JOIN categories ON (products.category_id = categories.id)
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
		&product.Category.ID,
		&product.Category.CategoryName,
		&product.Category.CreatedAt,
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

func (s *ProductStore) GetAllByName(ctx context.Context, productName string) ([]Product, error) {
	query := `
		SELECT products.id, product_name, description, price, stock, size, color, products.created_at, products.updated_at, categories.*
		FROM products
		JOIN categories ON (products.category_id = categories.id)
		WHERE product_name ILIKE $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	productName = "" + productName + "%"

	rows, err := s.db.QueryContext(ctx, query, productName)
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
			&product.Category.ID,
			&product.Category.CategoryName,
			&product.Category.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		productList = append(productList, product)

	}

	return productList, err
}

func (s *ProductStore) GetAllByCategoryID(ctx context.Context, categoryID int64) ([]Product, error) {
	query := `
		SELECT products.id, product_name, description, price, stock, size, color, products.created_at, products.updated_at, categories.*
		FROM products
		JOIN categories ON (products.category_id = categories.id)
		WHERE category_id = $1 
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, categoryID)
	if err != nil {
		return nil, err
	}

	rows.Close()

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
			&product.Category.ID,
			&product.Category.CategoryName,
			&product.Category.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		productList = append(productList, product)

	}

	return productList, nil
}

func (s *ProductStore) Delete(ctx context.Context, productID int64) error {
	query := `
		DELETE FROM products WHERE product_name = $1
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, productID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProductStore) Update(ctx context.Context, product *Product) error {
	query := `
		UPDATE products
		SET product_name = $2, description = $3, price = $4, stock = $5, size = $6,
		color = $7, category_id = $8
		WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		product.ID,
		product.ProductName,
		product.Description,
		product.Price,
		product.Stock,
		product.Size,
		product.Color,
		product.CategoryID,
	).Scan()
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrNotFound
		default:
			return err
		}
	}

	return nil
}
