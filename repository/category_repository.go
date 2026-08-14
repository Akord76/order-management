package repository

import (
	"order-management/model"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(category *model.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepository) FindAll() ([]model.Category, error) {
	var categories []model.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) FindByID(id int) (*model.Category, error) {
	var category model.Category
	if err := r.db.First(&category, "CategoryId = ?", id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) Update(category *model.Category) error {
	return r.db.Model(&model.Category{}).
		Where("CategoryId = ?", category.CategoryID).
		Updates(category).Error
}

func (r *CategoryRepository) Delete(id int) error {
	return r.db.Delete(&model.Category{}, "CategoryId = ?", id).Error
}
