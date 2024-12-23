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

func (s *OrderStore) Create(ctx context.Context, order *Order) error {
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
func (s *OrderStore) GetAll(ctx context.Context) ([]Order, error) {
	query := `
		SELECT id, user_id, total_price, status, created_at
		FROM orders
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orderList []Order

	for rows.Next() {
		var order Order
		err := rows.Scan(
			&order.ID,
			&order.UserID,
			&order.TotalPrice,
			&order.Status,
			&order.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		orderList = append(orderList, order)

	}

	return orderList, nil
}

func (s *OrderStore) GetAllByUserID(ctx context.Context, userID int64) ([]Order, error) {
	query := `
		SELECT id, user_id, total_price, status, created_at
		FROM orders
		WHERE user_id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)

	defer cancel()

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orderList []Order

	for rows.Next() {
		var order Order

		err := rows.Scan(
			&order.ID,
			&order.UserID,
			&order.TotalPrice,
			&order.Status,
			&order.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		orderList = append(orderList, order)

	}

	return orderList, nil
}

func (s *OrderStore) Update(ctx context.Context, order *Order) error {
	query := `
		UPDATE orders
		SET total_price = $2, status = $3
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)

	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		order.ID,
		order.TotalPrice,
		order.Status,
	).Scan()
	if err != nil {
		return err
	}

	return nil
}
