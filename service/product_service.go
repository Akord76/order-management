package service

import (
	"errors"

	"order-management/model"
	"order-management/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(product *model.Product) error {
	if product.ProductID == "" || product.ProductName == "" {
		return errors.New("product id and product name are required")
	}
	return s.repo.Create(product)
}

func (s *ProductService) GetAll() ([]model.Product, error) {
	return s.repo.FindAll()
}

func (s *ProductService) GetByID(productNumber int, productID string) (*model.Product, error) {
	return s.repo.FindByID(productNumber, productID)
}

func (s *ProductService) Update(product *model.Product) error {
	if product.ProductID == "" {
		return errors.New("product id is required")
	}
	return s.repo.Update(product)
}

func (s *ProductService) Delete(productNumber int, productID string) error {
	return s.repo.Delete(productNumber, productID)
}
