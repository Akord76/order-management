package service

import (
	"errors"

	"order-management/model"
	"order-management/repository"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) Create(order *model.OrderMaster) error {
	if order.OrderNo == "" {
		return errors.New("order no is required")
	}
	if len(order.Details) == 0 {
		return errors.New("order must contain at least one detail line")
	}
	return s.repo.CreateWithDetails(order)
}

func (s *OrderService) GetAll() ([]model.OrderMaster, error) {
	return s.repo.FindAll()
}

func (s *OrderService) GetByID(orderID int, orderNo string) (*model.OrderMaster, error) {
	return s.repo.FindByID(orderID, orderNo)
}

func (s *OrderService) Update(order *model.OrderMaster) error {
	if order.OrderNo == "" {
		return errors.New("order no is required")
	}
	return s.repo.Update(order)
}

func (s *OrderService) Delete(orderID int, orderNo string) error {
	return s.repo.Delete(orderID, orderNo)
}

func (s *OrderService) AddDetail(detail *model.OrderDetail) error {
	if detail.OrderNo == "" || detail.ItemName == "" {
		return errors.New("order no and item name are required")
	}
	return s.repo.AddDetail(detail)
}

func (s *OrderService) GetDetails(orderNo string) ([]model.OrderDetail, error) {
	return s.repo.FindDetailsByOrderNo(orderNo)
}

func (s *OrderService) UpdateDetail(detail *model.OrderDetail) error {
	return s.repo.UpdateDetail(detail)
}

func (s *OrderService) DeleteDetail(orderDetailNo int, orderNo string) error {
	return s.repo.DeleteDetail(orderDetailNo, orderNo)
}
