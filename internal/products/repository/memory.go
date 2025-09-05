package repository

import "github.com/akmyrzza/electrohub/internal/products/entity"

type InMemoryProductRepository struct {
	data []entity.Product
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{data: make([]entity.Product, 0)}
}

func (r *InMemoryProductRepository) ListProducts() ([]entity.Product, error) {
	return r.data, nil
}

func (r *InMemoryProductRepository) CreateProduct(p entity.Product) error {
	p.ID = int64(len(r.data) + 1)
	r.data = append(r.data, p)
	return nil
}
