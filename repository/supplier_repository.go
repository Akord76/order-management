package repository

import (
	"order-management/model"

	"gorm.io/gorm"
)

type SupplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) *SupplierRepository {
	return &SupplierRepository{db: db}
}

func (r *SupplierRepository) Create(supplier *model.Supplier) error {
	return r.db.Create(supplier).Error
}

func (r *SupplierRepository) FindAll() ([]model.Supplier, error) {
	var suppliers []model.Supplier
	err := r.db.Find(&suppliers).Error
	return suppliers, err
}

func (r *SupplierRepository) FindByID(supplierNumber int, supplierID string) (*model.Supplier, error) {
	var supplier model.Supplier
	err := r.db.Where("SuplierNumber = ? AND SuplierID = ?", supplierNumber, supplierID).
		First(&supplier).Error
	if err != nil {
		return nil, err
	}
	return &supplier, nil
}

func (r *SupplierRepository) FindBySupplierID(supplierID string) (*model.Supplier, error) {
	var supplier model.Supplier
	err := r.db.Where("SuplierID = ?", supplierID).First(&supplier).Error
	if err != nil {
		return nil, err
	}
	return &supplier, nil
}

func (r *SupplierRepository) Update(supplier *model.Supplier) error {
	return r.db.Model(&model.Supplier{}).
		Where("SuplierNumber = ? AND SuplierID = ?", supplier.SupplierNumber, supplier.SupplierID).
		Updates(supplier).Error
}

func (r *SupplierRepository) Delete(supplierNumber int, supplierID string) error {
	return r.db.Where("SuplierNumber = ? AND SuplierID = ?", supplierNumber, supplierID).
		Delete(&model.Supplier{}).Error
}
