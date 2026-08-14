package model

// CashReceipt maps to dbo.cashreceipt.
type CashReceipt struct {
	ChasNumber int     `gorm:"column:ChasNumber;primaryKey;autoIncrement" json:"chas_number"`
	Name       string  `gorm:"column:Name;size:20" json:"name"`
	ChasValue  float64 `gorm:"column:ChasValue" json:"chas_value"`
	Notes      string  `gorm:"column:Notes;size:100" json:"notes"`
}

func (CashReceipt) TableName() string { return "cashreceipt" }
