package model

import "time"

// OrderSaleMaster maps to dbo.OrderSaleMaster.
// Primary key is composite: OrderSaleNo + OrderSaleID.
type OrderSaleMaster struct {
	OrderSaleNo   int        `gorm:"column:OrderSaleNo;primaryKey" json:"order_sale_no"`
	OrderSaleID   string     `gorm:"column:OrderSaleID;primaryKey;size:50" json:"order_sale_id"`
	OrderSaleDate *time.Time `gorm:"column:OrderSaleDate;type:datetime" json:"order_sale_date"`
	CustomerID    string     `gorm:"column:CustomerID;size:50" json:"customer_id"`
	Shipment      string     `gorm:"column:Shipment;size:20" json:"shipment"`
	ShipNumber    string     `gorm:"column:ShipNumber;size:20" json:"ship_number"`
	DriverNumber  string     `gorm:"column:DriverNumber;size:20" json:"driver_number"`
	Description   string     `gorm:"column:Description;size:100" json:"description"`

	Details []OrderSaleDetail `gorm:"foreignKey:OrderSaleNo;references:OrderSaleNo" json:"details,omitempty"`
}

func (OrderSaleMaster) TableName() string { return "OrderSaleMaster" }
