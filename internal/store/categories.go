package store

import (
	"context"
	"database/sql"
)

type Category struct {
	ID           int64  `json:"id"`
	CategoryName string `json:"category_name"`
	CreatedAt    string `json:"created_at"`
}

type CategoryStore struct {
	db *sql.DB
}

func (s *CategoryStore) Create(ctx context.Context, category *Category) error {
	query := `
		INSERT INTO categories(category_name) VALUES ($1)
		RETURNING id, created_at
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		category.CategoryName,
	).Scan(
		&category.ID,
		&category.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *CategoryStore) GetAll(ctx context.Context) ([]Category, error) {
	query := `
		SELECT id, category_name 
		FROM categories
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []Category

	for rows.Next() {
		var category Category

		err := rows.Scan(
			&category.ID,
			&category.CategoryName,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
