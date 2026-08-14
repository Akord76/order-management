package service

import (
	"errors"

	"order-management/model"
	"order-management/repository"
)

type OrderSaleService struct {
	repo *repository.OrderSaleRepository
}

func NewOrderSaleService(repo *repository.OrderSaleRepository) *OrderSaleService {
	return &OrderSaleService{repo: repo}
}

func (s *OrderSaleService) Create(orderSale *model.OrderSaleMaster) error {
	if orderSale.OrderSaleID == "" {
		return errors.New("order sale id is required")
	}
	if len(orderSale.Details) == 0 {
		return errors.New("order sale must contain at least one detail line")
	}
	return s.repo.CreateWithDetails(orderSale)
}

func (s *OrderSaleService) GetAll() ([]model.OrderSaleMaster, error) {
	return s.repo.FindAll()
}

func (s *OrderSaleService) GetByID(orderSaleNo int, orderSaleID string) (*model.OrderSaleMaster, error) {
	return s.repo.FindByID(orderSaleNo, orderSaleID)
}

func (s *OrderSaleService) Update(orderSale *model.OrderSaleMaster) error {
	if orderSale.OrderSaleID == "" {
		return errors.New("order sale id is required")
	}
	return s.repo.Update(orderSale)
}

func (s *OrderSaleService) Delete(orderSaleNo int, orderSaleID string) error {
	return s.repo.Delete(orderSaleNo, orderSaleID)
}

func (s *OrderSaleService) AddDetail(detail *model.OrderSaleDetail) error {
	if detail.OrderSaleDetailID == "" || detail.ItemName == "" {
		return errors.New("order sale detail id and item name are required")
	}
	return s.repo.AddDetail(detail)
}

func (s *OrderSaleService) GetDetails(orderSaleNo int) ([]model.OrderSaleDetail, error) {
	return s.repo.FindDetailsByOrderSaleNo(orderSaleNo)
}

func (s *OrderSaleService) UpdateDetail(detail *model.OrderSaleDetail) error {
	return s.repo.UpdateDetail(detail)
}

func (s *OrderSaleService) DeleteDetail(orderSaleDetailNo int, orderSaleDetailID string) error {
	return s.repo.DeleteDetail(orderSaleDetailNo, orderSaleDetailID)
}
