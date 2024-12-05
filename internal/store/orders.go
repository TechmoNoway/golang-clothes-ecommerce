package store

import (
	"context"
	"database/sql"
)

type Order struct {
	ID         int64  `json:"id"`
	UserID     int64  `json:"user_id"`
	TotalPrice int64  `json:"total_price"`
	Status     string `json:"status"`
	CreatedAt  string `json:"created_at"`
}

type OrderStore struct {
	db *sql.DB
}

func (s *OrderStore) create(ctx context.Context, order *Order) error {
	query := `INSERT INTO orders (user_id, total_price, status) VALUES ($1, $2, $3)`

	err := s.db.QueryRowContext(
		ctx,
		query,
		order.UserID,
		order.TotalPrice,
		order.Status,
	).Scan(
		&order.ID,
		&order.CreatedAt,
	)
	if err != nil {
		return nil
	}

	return nil
}
