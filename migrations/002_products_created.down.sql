ALTER TABLE products
    DROP CONSTRAINT IF EXISTS products_name_unique;

ALTER TABLE products
    DROP COLUMN IF EXISTS created_at,
    DROP COLUMN IF EXISTS updated_at;

DROP INDEX IF EXISTS idx_products_name;