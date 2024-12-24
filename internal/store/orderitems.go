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

func (s *OrderItemStore) Create(ctx context.Context, orderItem *OrderItem) error {

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

func (s *OrderItemStore) GetAll(ctx context.Context) ([]OrderItem, error) {
	query := `
		SELECT id, order_id, product_id, quantity, price, created_at
		FROM order_items

	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)

	defer cancel()

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	rows.Close()

	var orderItemList []OrderItem

	for rows.Next() {
		var orderItem OrderItem

		err := rows.Scan(
			&orderItem.ID,
			&orderItem.OrderID,
			&orderItem.ProductID,
			&orderItem.Quantity,
			&orderItem.Price,
			&orderItem.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		orderItemList = append(orderItemList, orderItem)

	}

	return orderItemList, err
}
