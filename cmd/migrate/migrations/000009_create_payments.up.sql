CREATE TABLE payments (
    id BIGSERIAL PRIMARY KEY,
    order_id BIGINT REFERENCES orders(id),
    payment_method TEXT,
    status TEXT,
    payment_date TIMESTAMP WITH TIME ZONE
);