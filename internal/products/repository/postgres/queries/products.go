package queries

const (
	InsertProduct = `
		INSERT INTO products (name, price, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	GetProduct = `
		SELECT id, name, price, created_at, updated_at
		FROM products WHERE id = $1
	`

	UpdateProduct = `
		UPDATE products
		SET name = $1, price = $2, updated_at = $3
		WHERE id = $4
		RETURNING id, name, price, created_at, updated_at
	`

	DeleteProduct = `
		DELETE FROM products WHERE id = $1
	`

	GetProducts = `
		SELECT id, name, price, created_at, updated_at
		FROM products
	`
)
