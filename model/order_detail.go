package model

// OrderDetail maps to dbo.OrderDetail.
// Primary key is composite: OrderDetailNo (identity) + OrderNo.
type OrderDetail struct {
	OrderDetailNo int     `gorm:"column:OrderDetailNo;primaryKey;autoIncrement" json:"order_detail_no"`
	OrderNo       string  `gorm:"column:OrderNo;primaryKey;size:20" json:"order_no"`
	ItemName      string  `gorm:"column:ItemName;size:20" json:"item_name"`
	Measure       string  `gorm:"column:Measure;size:20" json:"measure"`
	Qty           int     `gorm:"column:Qty" json:"qty"`
	Price         float64 `gorm:"column:Price" json:"price"`
}

func (OrderDetail) TableName() string { return "OrderDetail" }
