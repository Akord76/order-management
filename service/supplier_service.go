package service

import (
	"errors"

	"order-management/model"
	"order-management/repository"
)

type SupplierService struct {
	repo *repository.SupplierRepository
}

func NewSupplierService(repo *repository.SupplierRepository) *SupplierService {
	return &SupplierService{repo: repo}
}

func (s *SupplierService) Create(supplier *model.Supplier) error {
	if supplier.SupplierID == "" || supplier.SupplierName == "" {
		return errors.New("supplier id and supplier name are required")
	}
	return s.repo.Create(supplier)
}

func (s *SupplierService) GetAll() ([]model.Supplier, error) {
	return s.repo.FindAll()
}

func (s *SupplierService) GetByID(supplierNumber int, supplierID string) (*model.Supplier, error) {
	return s.repo.FindByID(supplierNumber, supplierID)
}

func (s *SupplierService) Update(supplier *model.Supplier) error {
	if supplier.SupplierID == "" {
		return errors.New("supplier id is required")
	}
	return s.repo.Update(supplier)
}

func (s *SupplierService) Delete(supplierNumber int, supplierID string) error {
	return s.repo.Delete(supplierNumber, supplierID)
}
