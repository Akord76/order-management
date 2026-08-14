package repository

import (
	"order-management/model"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) Create(customer *model.Customer) error {
	return r.db.Create(customer).Error
}

func (r *CustomerRepository) FindAll() ([]model.Customer, error) {
	var customers []model.Customer
	err := r.db.Find(&customers).Error
	return customers, err
}

func (r *CustomerRepository) FindByID(custNumber int, custID string) (*model.Customer, error) {
	var customer model.Customer
	err := r.db.Where("CustNumber = ? AND CustID = ?", custNumber, custID).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerRepository) FindByCustID(custID string) (*model.Customer, error) {
	var customer model.Customer
	err := r.db.Where("CustID = ?", custID).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerRepository) Update(customer *model.Customer) error {
	return r.db.Model(&model.Customer{}).
		Where("CustNumber = ? AND CustID = ?", customer.CustNumber, customer.CustID).
		Updates(customer).Error
}

func (r *CustomerRepository) Delete(custNumber int, custID string) error {
	return r.db.Where("CustNumber = ? AND CustID = ?", custNumber, custID).
		Delete(&model.Customer{}).Error
}
