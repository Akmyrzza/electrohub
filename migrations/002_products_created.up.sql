ALTER TABLE products
    ADD COLUMN created_at TIMESTAMP DEFAULT now(),
    ADD COLUMN updated_at TIMESTAMP DEFAULT now();

ALTER TABLE products
    ADD CONSTRAINT products_name_unique UNIQUE (name);

CREATE INDEX IF NOT EXISTS idx_products_name ON products(name);