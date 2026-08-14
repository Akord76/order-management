package model

// CommitmentFee maps to dbo.CommitmentFee.
type CommitmentFee struct {
	CommitmentID       int     `gorm:"column:CommitmentID;primaryKey;autoIncrement" json:"commitment_id"`
	EmployeeCardNumber *int    `gorm:"column:EmployeeCardNumber" json:"employee_card_number"`
	CommitmentValue    float64 `gorm:"column:CommitmentValue" json:"commitment_value"`
	ParameterFee       *int    `gorm:"column:ParameterFee" json:"parameter_fee"`
	ProductID          string  `gorm:"column:ProductID;size:25" json:"product_id"`
}

func (CommitmentFee) TableName() string { return "CommitmentFee" }
