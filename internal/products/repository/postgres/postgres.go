package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/akmyrzza/electrohub/internal/products/entity"
	"github.com/akmyrzza/electrohub/internal/products/repository/postgres/queries"
)

type PostgresProductRepository struct {
	db *sql.DB
}

func NewPostgresProductRepository(db *sql.DB) *PostgresProductRepository {
	return &PostgresProductRepository{db: db}
}

func (r *PostgresProductRepository) CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error) {
	var id int64
	row := r.db.QueryRowContext(ctx, queries.InsertProduct, product.Name, product.Price, time.Now().UTC(), time.Now().UTC())
	err := row.Scan(&id)
	if err != nil {
		return entity.Product{}, err
	}

	product.ID = id
	return product, nil
}

func (r *PostgresProductRepository) GetProductByID(ctx context.Context, id int64) (entity.Product, error) {
	var product entity.Product
	row := r.db.QueryRowContext(ctx, queries.GetProduct, id)
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return entity.Product{}, err
	}

	return product, nil
}

func (r *PostgresProductRepository) UpdateProduct(ctx context.Context, id int64, product entity.Product) (entity.Product, error) {
	var resultProduct entity.Product
	row := r.db.QueryRowContext(ctx, queries.UpdateProduct, product.Name, product.Price, time.Now().UTC(), id)
	err := row.Scan(&resultProduct.ID, &resultProduct.Name, &resultProduct.Price, &resultProduct.CreatedAt, &resultProduct.UpdatedAt)
	if err != nil {
		return entity.Product{}, err
	}

	return resultProduct, nil
}

func (r *PostgresProductRepository) DeleteProduct(ctx context.Context, id int64) error {
	result, err := r.db.ExecContext(ctx, queries.DeleteProduct, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (r *PostgresProductRepository) ListProducts(ctx context.Context) ([]entity.Product, error) {
	rows, err := r.db.QueryContext(ctx, queries.GetProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entity.Product
	for rows.Next() {
		var product entity.Product
		errScan := rows.Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		if errScan != nil {
			return nil, errScan
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
