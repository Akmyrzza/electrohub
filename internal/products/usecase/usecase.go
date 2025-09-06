package usecase

import (
	"context"

	"github.com/akmyrzza/electrohub/internal/products/entity"
)

type ProductRepository interface {
	ListProducts(ctx context.Context) ([]entity.Product, error)
	GetProductByID(ctx context.Context, id int64) (entity.Product, error)
	CreateProduct(ctx context.Context, p entity.Product) (entity.Product, error)
	UpdateProduct(ctx context.Context, id int64, p entity.Product) (entity.Product, error)
	DeleteProduct(ctx context.Context, id int64) error
}

type ProductService struct {
	repo ProductRepository
}

func NewProductService(r ProductRepository) *ProductService {
	return &ProductService{repo: r}
}

func (s *ProductService) ListProducts(ctx context.Context) ([]entity.Product, error) {
	return s.repo.ListProducts(ctx)
}

func (s *ProductService) GetProductByID(ctx context.Context, id int64) (entity.Product, error) {
	return s.repo.GetProductByID(ctx, id)
}

func (s *ProductService) CreateProduct(ctx context.Context, p entity.Product) (entity.Product, error) {
	return s.repo.CreateProduct(ctx, p)
}

func (s *ProductService) UpdateProduct(ctx context.Context, id int64, p entity.Product) (entity.Product, error) {
	return s.repo.UpdateProduct(ctx, id, p)
}

func (s *ProductService) DeleteProduct(ctx context.Context, id int64) error {
	return s.repo.DeleteProduct(ctx, id)
}
