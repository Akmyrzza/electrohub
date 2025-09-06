package repository

import (
	"errors"

	"github.com/akmyrzza/electrohub/internal/products/entity"
)

type InMemoryProductRepository struct {
	data []entity.Product
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{data: make([]entity.Product, 0)}
}

func (r *InMemoryProductRepository) ListProducts() ([]entity.Product, error) {
	return r.data, nil
}

func (r *InMemoryProductRepository) GetProductByID(id int64) (entity.Product, error) {
	for _, product := range r.data {
		if product.ID == id {
			return product, nil
		}
	}

	return entity.Product{}, errors.New("product not found")
}

func (r *InMemoryProductRepository) CreateProduct(p entity.Product) (entity.Product, error) {
	p.ID = int64(len(r.data) + 1)
	r.data = append(r.data, p)
	return p, nil
}

func (r *InMemoryProductRepository) UpdateProduct(id int64, updated entity.Product) (entity.Product, error) {
	for i, p := range r.data {
		if p.ID == id {
			r.data[i] = updated
			return r.data[i], nil
		}
	}
	return entity.Product{}, errors.New("product not found")
}

func (r *InMemoryProductRepository) DeleteProduct(id int64) error {
	for i, p := range r.data {
		if p.ID == id {
			r.data = append(r.data[:i], r.data[i+1:]...)
			return nil
		}
	}
	return errors.New("product not found")
}
