package repository

import (
	"order-management/model"

	"gorm.io/gorm"
)

type OrderSaleRepository struct {
	db *gorm.DB
}

func NewOrderSaleRepository(db *gorm.DB) *OrderSaleRepository {
	return &OrderSaleRepository{db: db}
}

// CreateWithDetails inserts an OrderSaleMaster together with its
// OrderSaleDetail rows inside a single transaction.
func (r *OrderSaleRepository) CreateWithDetails(orderSale *model.OrderSaleMaster) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(orderSale).Error; err != nil {
			return err
		}
		for i := range orderSale.Details {
			orderSale.Details[i].OrderSaleNo = orderSale.OrderSaleNo
			if err := tx.Create(&orderSale.Details[i]).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *OrderSaleRepository) FindAll() ([]model.OrderSaleMaster, error) {
	var orderSales []model.OrderSaleMaster
	err := r.db.Preload("Details").Find(&orderSales).Error
	return orderSales, err
}

func (r *OrderSaleRepository) FindByID(orderSaleNo int, orderSaleID string) (*model.OrderSaleMaster, error) {
	var orderSale model.OrderSaleMaster
	err := r.db.Preload("Details").
		Where("OrderSaleNo = ? AND OrderSaleID = ?", orderSaleNo, orderSaleID).
		First(&orderSale).Error
	if err != nil {
		return nil, err
	}
	return &orderSale, nil
}

func (r *OrderSaleRepository) FindByOrderSaleNo(orderSaleNo int) (*model.OrderSaleMaster, error) {
	var orderSale model.OrderSaleMaster
	err := r.db.Preload("Details").Where("OrderSaleNo = ?", orderSaleNo).First(&orderSale).Error
	if err != nil {
		return nil, err
	}
	return &orderSale, nil
}

func (r *OrderSaleRepository) Update(orderSale *model.OrderSaleMaster) error {
	return r.db.Model(&model.OrderSaleMaster{}).
		Where("OrderSaleNo = ? AND OrderSaleID = ?", orderSale.OrderSaleNo, orderSale.OrderSaleID).
		Updates(orderSale).Error
}

func (r *OrderSaleRepository) Delete(orderSaleNo int, orderSaleID string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("OrderSaleNo = ?", orderSaleNo).Delete(&model.OrderSaleDetail{}).Error; err != nil {
			return err
		}
		return tx.Where("OrderSaleNo = ? AND OrderSaleID = ?", orderSaleNo, orderSaleID).
			Delete(&model.OrderSaleMaster{}).Error
	})
}

// --- OrderSaleDetail specific helpers ---

func (r *OrderSaleRepository) AddDetail(detail *model.OrderSaleDetail) error {
	return r.db.Create(detail).Error
}

func (r *OrderSaleRepository) FindDetailsByOrderSaleNo(orderSaleNo int) ([]model.OrderSaleDetail, error) {
	var details []model.OrderSaleDetail
	err := r.db.Where("OrderSaleNo = ?", orderSaleNo).Find(&details).Error
	return details, err
}

func (r *OrderSaleRepository) UpdateDetail(detail *model.OrderSaleDetail) error {
	return r.db.Model(&model.OrderSaleDetail{}).
		Where("OrderSaleDetailNo = ? AND OrderSaleDetailID = ?", detail.OrderSaleDetailNo, detail.OrderSaleDetailID).
		Updates(detail).Error
}

func (r *OrderSaleRepository) DeleteDetail(orderSaleDetailNo int, orderSaleDetailID string) error {
	return r.db.Where("OrderSaleDetailNo = ? AND OrderSaleDetailID = ?", orderSaleDetailNo, orderSaleDetailID).
		Delete(&model.OrderSaleDetail{}).Error
}
