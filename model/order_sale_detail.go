package model

// OrderSaleDetail maps to dbo.OrderSaleDetail.
// Primary key is composite: OrderSaleDetailNo (identity) + OrderSaleDetailID.
// OrderSaleNo links back to OrderSaleMaster.OrderSaleNo.
type OrderSaleDetail struct {
	OrderSaleDetailNo int     `gorm:"column:OrderSaleDetailNo;primaryKey;autoIncrement" json:"order_sale_detail_no"`
	OrderSaleDetailID string  `gorm:"column:OrderSaleDetailID;primaryKey;size:20" json:"order_sale_detail_id"`
	OrderSaleNo       int     `gorm:"column:OrderSaleNo;not null" json:"order_sale_no"`
	ItemName          string  `gorm:"column:ItemName;size:20" json:"item_name"`
	Measure           string  `gorm:"column:Measure;size:20" json:"measure"`
	Qty               int     `gorm:"column:Qty" json:"qty"`
	Price             float64 `gorm:"column:Price" json:"price"`
	DocumentNumber    string  `gorm:"column:DokumenNumber;size:25" json:"document_number"`
}

func (OrderSaleDetail) TableName() string { return "OrderSaleDetail" }
