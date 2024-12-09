ALTER TABLE
    order_items
ADD CONSTRAINT fk_order_item_order FOREIGN KEY (order_id) REFERENCES orders(id);
ALTER TABLE
    order_items
ADD CONSTRAINT fk_order_item_product FOREIGN KEY (product_id) REFERENCES products(id);
