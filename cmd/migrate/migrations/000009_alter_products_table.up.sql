ALTER TABLE
    products
ADD CONSTRAINT fk_product_category FOREIGN KEY (category_id) REFERENCES categories(id);