package model

// SharingProfit maps to dbo.SharingProvit.
// Primary key is composite: SharingProvitNum + SharingProvitID.
type SharingProfit struct {
	SharingProfitNum int     `gorm:"column:SharingProvitNum;primaryKey" json:"sharing_profit_num"`
	SharingProfitID  string  `gorm:"column:SharingProvitID;primaryKey;size:15" json:"sharing_profit_id"`
	CommitmentID     *int    `gorm:"column:CommitmentID" json:"commitment_id"`
	OrderSaleDetailID string `gorm:"column:OrderSaleDetailID;size:20" json:"order_sale_detail_id"`
	ShareValue       float64 `gorm:"column:ShareValue" json:"share_value"`
}

func (SharingProfit) TableName() string { return "SharingProvit" }
