package store

import (
	"context"
	"database/sql"
)

type Payment struct {
	ID            int64  `json:"id"`
	OrderID       int64  `json:"order_id"`
	PaymentMethod string `json:"payment_method"`
	Status        string `json:"status"`
	PaymentDate   string `json:"payment_date"`
}

type PaymentStore struct {
	db *sql.DB
}

func (s *PaymentStore) create(ctx context.Context, payment Payment) error {
	query := `
		INSERT INTO payments (order_id, paymetn_method, status)
		VALUES ($1, $2, $3, $4)
		`
	err := s.db.QueryRowContext(
		ctx,
		query,
		payment.OrderID,
		payment.PaymentMethod,
		payment.Status,
	).Scan(
		&payment.ID,
		&payment.PaymentDate,
	)
	if err != nil {
		return err
	}

	return nil
}
