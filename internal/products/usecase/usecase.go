package usecase

import "github.com/akmyrzza/electrohub/internal/products/entity"

type ProductRepository interface {
	ListProducts() ([]entity.Product, error)
	CreateProduct(p entity.Product) error
}

type ProductService struct {
	repo ProductRepository
}

func NewProductService(r ProductRepository) *ProductService {
	return &ProductService{repo: r}
}

func (s *ProductService) ListProducts() ([]entity.Product, error) {
	return s.repo.ListProducts()
}

func (s *ProductService) CreateProduct(p entity.Product) error {
	return s.repo.CreateProduct(p)
}
