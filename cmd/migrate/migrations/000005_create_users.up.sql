CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password BYTEA NOT NULL, 
    avatar_url VARCHAR(255),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    phone VARCHAR(10) NOT NULL,
    address VARCHAR(255) NOT NULL,
    role_id BIGINT REFERENCES roles(id), exists
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);