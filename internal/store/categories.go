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

type CategoryStorage struct {
	db *sql.DB
}

func (s *CategoryStorage) create(ctx context.Context, category *Category) error {
	query := `
		INSERT INTO categories(category_name) VALUES ($1)
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
