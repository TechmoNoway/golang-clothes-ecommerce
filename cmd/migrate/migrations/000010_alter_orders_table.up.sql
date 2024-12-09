ALTER TABLE
    orders
ADD CONSTRAINT fk_order_user FOREIGN KEY (user_id) REFERENCES users(id);
