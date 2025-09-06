package usecase

import "github.com/akmyrzza/electrohub/internal/products/entity"

type ProductRepository interface {
	ListProducts() ([]entity.Product, error)
	GetProductByID(id int64) (entity.Product, error)
	CreateProduct(p entity.Product) (entity.Product, error)
	UpdateProduct(id int64, p entity.Product) (entity.Product, error)
	DeleteProduct(id int64) error
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

func (s *ProductService) GetProductByID(id int64) (entity.Product, error) {
	return s.repo.GetProductByID(id)
}

func (s *ProductService) CreateProduct(p entity.Product) (entity.Product, error) {
	return s.repo.CreateProduct(p)
}

func (s *ProductService) UpdateProduct(id int64, p entity.Product) (entity.Product, error) {
	return s.repo.UpdateProduct(id, p)
}

func (s *ProductService) DeleteProduct(id int64) error {
	return s.repo.DeleteProduct(id)
}
