package repository

import (
	"order-management/model"

	"gorm.io/gorm"
)

type CommitmentFeeRepository struct {
	db *gorm.DB
}

func NewCommitmentFeeRepository(db *gorm.DB) *CommitmentFeeRepository {
	return &CommitmentFeeRepository{db: db}
}

func (r *CommitmentFeeRepository) Create(fee *model.CommitmentFee) error {
	return r.db.Create(fee).Error
}

func (r *CommitmentFeeRepository) FindAll() ([]model.CommitmentFee, error) {
	var fees []model.CommitmentFee
	err := r.db.Find(&fees).Error
	return fees, err
}

func (r *CommitmentFeeRepository) FindByID(id int) (*model.CommitmentFee, error) {
	var fee model.CommitmentFee
	if err := r.db.First(&fee, "CommitmentID = ?", id).Error; err != nil {
		return nil, err
	}
	return &fee, nil
}

func (r *CommitmentFeeRepository) FindByEmployeeCard(cardNumber int) ([]model.CommitmentFee, error) {
	var fees []model.CommitmentFee
	err := r.db.Where("EmployeeCardNumber = ?", cardNumber).Find(&fees).Error
	return fees, err
}

func (r *CommitmentFeeRepository) Update(fee *model.CommitmentFee) error {
	return r.db.Model(&model.CommitmentFee{}).
		Where("CommitmentID = ?", fee.CommitmentID).
		Updates(fee).Error
}

func (r *CommitmentFeeRepository) Delete(id int) error {
	return r.db.Delete(&model.CommitmentFee{}, "CommitmentID = ?", id).Error
}
