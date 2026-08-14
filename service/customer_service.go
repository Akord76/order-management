package service

import (
	"errors"

	"order-management/model"
	"order-management/repository"
)

type CustomerService struct {
	repo *repository.CustomerRepository
}

func NewCustomerService(repo *repository.CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) Create(customer *model.Customer) error {
	if customer.CustID == "" || customer.CustName == "" {
		return errors.New("customer id and customer name are required")
	}
	return s.repo.Create(customer)
}

func (s *CustomerService) GetAll() ([]model.Customer, error) {
	return s.repo.FindAll()
}

func (s *CustomerService) GetByID(custNumber int, custID string) (*model.Customer, error) {
	return s.repo.FindByID(custNumber, custID)
}

func (s *CustomerService) Update(customer *model.Customer) error {
	if customer.CustID == "" {
		return errors.New("customer id is required")
	}
	return s.repo.Update(customer)
}

func (s *CustomerService) Delete(custNumber int, custID string) error {
	return s.repo.Delete(custNumber, custID)
}
