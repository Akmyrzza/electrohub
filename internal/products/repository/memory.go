package repository

import "github.com/akmyrzza/electrohub/internal/products/entity"

type InMemoryProductRepository struct {
	data []entity.Product
}

func NewInMemoryProductRepo() *InMemoryProductRepository {
	return &InMemoryProductRepository{data: []entity.Product{}}
}

func (r *InMemoryProductRepository) GetAll() ([]entity.Product, error) {
	return r.data, nil
}

func (r *InMemoryProductRepository) Create(p entity.Product) error {
	p.ID = int64(len(r.data) + 1)
	r.data = append(r.data, p)
	return nil
}
