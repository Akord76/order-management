package repository

import (
	"order-management/model"

	"gorm.io/gorm"
)

type CashReceiptRepository struct {
	db *gorm.DB
}

func NewCashReceiptRepository(db *gorm.DB) *CashReceiptRepository {
	return &CashReceiptRepository{db: db}
}

func (r *CashReceiptRepository) Create(receipt *model.CashReceipt) error {
	return r.db.Create(receipt).Error
}

func (r *CashReceiptRepository) FindAll() ([]model.CashReceipt, error) {
	var receipts []model.CashReceipt
	err := r.db.Find(&receipts).Error
	return receipts, err
}

func (r *CashReceiptRepository) FindByID(id int) (*model.CashReceipt, error) {
	var receipt model.CashReceipt
	if err := r.db.First(&receipt, "ChasNumber = ?", id).Error; err != nil {
		return nil, err
	}
	return &receipt, nil
}

func (r *CashReceiptRepository) Update(receipt *model.CashReceipt) error {
	return r.db.Model(&model.CashReceipt{}).
		Where("ChasNumber = ?", receipt.ChasNumber).
		Updates(receipt).Error
}

func (r *CashReceiptRepository) Delete(id int) error {
	return r.db.Delete(&model.CashReceipt{}, "ChasNumber = ?", id).Error
}
