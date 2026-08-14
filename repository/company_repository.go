package repository

import (
	"order-management/model"

	"gorm.io/gorm"
)

type CompanyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{db: db}
}

func (r *CompanyRepository) Create(company *model.Company) error {
	return r.db.Create(company).Error
}

func (r *CompanyRepository) FindAll() ([]model.Company, error) {
	var companies []model.Company
	err := r.db.Find(&companies).Error
	return companies, err
}

func (r *CompanyRepository) FindByID(companyNumber int, companyID string) (*model.Company, error) {
	var company model.Company
	err := r.db.Where("CompanyNumber = ? AND CompanyID = ?", companyNumber, companyID).
		First(&company).Error
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *CompanyRepository) Update(company *model.Company) error {
	return r.db.Model(&model.Company{}).
		Where("CompanyNumber = ? AND CompanyID = ?", company.CompanyNumber, company.CompanyID).
		Updates(company).Error
}

func (r *CompanyRepository) Delete(companyNumber int, companyID string) error {
	return r.db.Where("CompanyNumber = ? AND CompanyID = ?", companyNumber, companyID).
		Delete(&model.Company{}).Error
}
