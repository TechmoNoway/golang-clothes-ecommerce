package store

import (
	"context"
	"database/sql"
)

type OrderItem struct {
	ID        int64  `json:"id"`
	OrderID   int64  `json:"order_id"`
	ProductID int64  `json:"product_id"`
	Quantity  int64  `json:"quantity"`
	Price     int64  `json:"price"`
	CreatedAt string `json:"created_at"`
}

type OrderItemStore struct {
	db *sql.DB
}

func (s *OrderItemStore) create(ctx context.Context, orderItem *OrderItem) error {

	query := `
        INSERT INTO order_items (order_id, product_id, quantity, price)
        VALUES ($1, $2, $3, $4)  
    `
	err := s.db.QueryRowContext(
		ctx,
		query,
		orderItem.OrderID,
		orderItem.ProductID,
		orderItem.Quantity,
		orderItem.Price,
	).Scan(
		&orderItem.ID,
		&orderItem.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
