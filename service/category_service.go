package service

import (
	"errors"

	"order-management/model"
	"order-management/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) Create(category *model.Category) error {
	if category.CategoryName == "" {
		return errors.New("category name is required")
	}
	return s.repo.Create(category)
}

func (s *CategoryService) GetAll() ([]model.Category, error) {
	return s.repo.FindAll()
}

func (s *CategoryService) GetByID(id int) (*model.Category, error) {
	return s.repo.FindByID(id)
}

func (s *CategoryService) Update(category *model.Category) error {
	if category.CategoryID == 0 {
		return errors.New("category id is required")
	}
	return s.repo.Update(category)
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
