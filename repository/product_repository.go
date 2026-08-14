package repository

import (
	"order-management/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) FindAll() ([]model.Product, error) {
	var products []model.Product
	err := r.db.Preload("Category").Find(&products).Error
	return products, err
}

// FindByID looks up a product by its composite key.
func (r *ProductRepository) FindByID(productNumber int, productID string) (*model.Product, error) {
	var product model.Product
	err := r.db.Preload("Category").
		Where("ProductNumber = ? AND ProductID = ?", productNumber, productID).
		First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// FindByProductID looks up a product by its unique ProductID only.
func (r *ProductRepository) FindByProductID(productID string) (*model.Product, error) {
	var product model.Product
	err := r.db.Preload("Category").
		Where("ProductID = ?", productID).
		First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) Update(product *model.Product) error {
	return r.db.Model(&model.Product{}).
		Where("ProductNumber = ? AND ProductID = ?", product.ProductNumber, product.ProductID).
		Updates(product).Error
}

func (r *ProductRepository) Delete(productNumber int, productID string) error {
	return r.db.Where("ProductNumber = ? AND ProductID = ?", productNumber, productID).
		Delete(&model.Product{}).Error
}
