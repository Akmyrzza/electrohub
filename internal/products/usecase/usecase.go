package usecase

import "github.com/akmyrzza/electrohub/internal/products/entity"

type ProductRepository interface {
	GetAll() ([]entity.Product, error)
	Create(p entity.Product) error
}

type ProductUseCase struct {
	repository ProductRepository
}

func NewProductUseCase(r ProductRepository) *ProductUseCase {
	return &ProductUseCase{repository: r}
}

func (uc *ProductUseCase) GetAll() ([]entity.Product, error) {
	return uc.repository.GetAll()
}

func (uc *ProductUseCase) Create(p entity.Product) error {
	return uc.repository.Create(p)
}
