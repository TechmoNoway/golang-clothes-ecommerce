CREATE TABLE payments (
    id BIGSERIAL PRIMARY KEY,
    order_id BIGINT,
    payment_method VARCHAR(255),
    status VARCHAR(255),
    payment_date TIMESTAMP WITH TIME ZONE
);