CREATE TABLE orders (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    total_price INTEGER NOT NULL,
    status VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);