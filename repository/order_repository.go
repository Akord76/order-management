package repository

import (
	"order-management/model"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// CreateWithDetails inserts an OrderMaster together with its OrderDetail
// rows inside a single transaction.
//
// Two things matter here, both guarding against the same class of bug
// (duplicate/extra OrderDetail rows):
//  1. tx.Omit("Details").Create(order) - stops GORM's automatic association
//     save from cascading an extra insert of order.Details when the Master
//     is created (see comment below).
//  2. Details are inserted as a single batch (tx.Create(&order.Details))
//     rather than one-by-one in a loop, which is the well-supported path
//     for GORM to correctly retrieve each row's generated identity
//     (OrderDetailNo is part of a composite primary key + SQL Server
//     IDENTITY column).
func (r *OrderRepository) CreateWithDetails(order *model.OrderMaster) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Omit("Details") is critical: OrderMaster.Details is tagged as a
		// GORM association (foreignKey:OrderNo). Without Omit, tx.Create(order)
		// would ALSO auto-cascade-insert order.Details itself (GORM's default
		// "save associations" behavior on Create), and then our explicit
		// batch insert below would insert the same rows a second time -
		// this was the actual cause of both the duplicate-key error and the
		// "2 rows became 4" bug. Omit makes this function the single place
		// that inserts OrderDetail rows.
		if err := tx.Omit("Details").Create(order).Error; err != nil {
			return err
		}

		for i := range order.Details {
			order.Details[i].OrderNo = order.OrderNo
			order.Details[i].OrderDetailNo = 0 // let the DB assign it; never trust a caller-supplied value here
		}

		if len(order.Details) > 0 {
			if err := tx.Create(&order.Details).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *OrderRepository) FindAll() ([]model.OrderMaster, error) {
	var orders []model.OrderMaster
	err := r.db.Preload("Details").Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) FindByID(orderID int, orderNo string) (*model.OrderMaster, error) {
	var order model.OrderMaster
	err := r.db.Preload("Details").
		Where("OrderID = ? AND OrderNo = ?", orderID, orderNo).
		First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) FindByOrderNo(orderNo string) (*model.OrderMaster, error) {
	var order model.OrderMaster
	err := r.db.Preload("Details").Where("OrderNo = ?", orderNo).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) Update(order *model.OrderMaster) error {
	return r.db.Model(&model.OrderMaster{}).
		Where("OrderID = ? AND OrderNo = ?", order.OrderID, order.OrderNo).
		Updates(order).Error
}

func (r *OrderRepository) Delete(orderID int, orderNo string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("OrderNo = ?", orderNo).Delete(&model.OrderDetail{}).Error; err != nil {
			return err
		}
		return tx.Where("OrderID = ? AND OrderNo = ?", orderID, orderNo).
			Delete(&model.OrderMaster{}).Error
	})
}

// --- OrderDetail specific helpers ---

func (r *OrderRepository) AddDetail(detail *model.OrderDetail) error {
	return r.db.Create(detail).Error
}

func (r *OrderRepository) FindDetailsByOrderNo(orderNo string) ([]model.OrderDetail, error) {
	var details []model.OrderDetail
	err := r.db.Where("OrderNo = ?", orderNo).Find(&details).Error
	return details, err
}

func (r *OrderRepository) UpdateDetail(detail *model.OrderDetail) error {
	return r.db.Model(&model.OrderDetail{}).
		Where("OrderDetailNo = ? AND OrderNo = ?", detail.OrderDetailNo, detail.OrderNo).
		Updates(detail).Error
}

func (r *OrderRepository) DeleteDetail(orderDetailNo int, orderNo string) error {
	return r.db.Where("OrderDetailNo = ? AND OrderNo = ?", orderDetailNo, orderNo).
		Delete(&model.OrderDetail{}).Error
}
