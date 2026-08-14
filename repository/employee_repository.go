package repository

import (
	"order-management/model"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) Create(employee *model.Employee) error {
	return r.db.Create(employee).Error
}

func (r *EmployeeRepository) FindAll() ([]model.Employee, error) {
	var employees []model.Employee
	err := r.db.Find(&employees).Error
	return employees, err
}

func (r *EmployeeRepository) FindByID(employeeID int, cardNumber int) (*model.Employee, error) {
	var employee model.Employee
	err := r.db.Where("EmployeeId = ? AND EmployeeCardNumber = ?", employeeID, cardNumber).
		First(&employee).Error
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *EmployeeRepository) FindByCardNumber(cardNumber int) (*model.Employee, error) {
	var employee model.Employee
	err := r.db.Where("EmployeeCardNumber = ?", cardNumber).First(&employee).Error
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *EmployeeRepository) Update(employee *model.Employee) error {
	return r.db.Model(&model.Employee{}).
		Where("EmployeeId = ? AND EmployeeCardNumber = ?", employee.EmployeeID, employee.EmployeeCardNumber).
		Updates(employee).Error
}

func (r *EmployeeRepository) Delete(employeeID int, cardNumber int) error {
	return r.db.Where("EmployeeId = ? AND EmployeeCardNumber = ?", employeeID, cardNumber).
		Delete(&model.Employee{}).Error
}
