package service

import (
	"errors"

	"order-management/model"
	"order-management/repository"
)

type EmployeeService struct {
	repo *repository.EmployeeRepository
}

func NewEmployeeService(repo *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) Create(employee *model.Employee) error {
	if employee.FirstName == "" || employee.EmployeeCardNumber == 0 {
		return errors.New("first name and employee card number are required")
	}
	return s.repo.Create(employee)
}

func (s *EmployeeService) GetAll() ([]model.Employee, error) {
	return s.repo.FindAll()
}

func (s *EmployeeService) GetByID(employeeID, cardNumber int) (*model.Employee, error) {
	return s.repo.FindByID(employeeID, cardNumber)
}

func (s *EmployeeService) Update(employee *model.Employee) error {
	if employee.EmployeeID == 0 {
		return errors.New("employee id is required")
	}
	return s.repo.Update(employee)
}

func (s *EmployeeService) Delete(employeeID, cardNumber int) error {
	return s.repo.Delete(employeeID, cardNumber)
}
