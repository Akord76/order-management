package model

import "time"

// OrderMaster maps to dbo.OrderMaster.
// Primary key is composite: OrderID + OrderNo.
type OrderMaster struct {
	OrderID        int        `gorm:"column:OrderID;primaryKey" json:"order_id"`
	OrderNo        string     `gorm:"column:OrderNo;primaryKey;size:20" json:"order_no"`
	OrderDate      *time.Time `gorm:"column:OrderDate;type:datetime" json:"order_date"`
	SupplierFrom   string     `gorm:"column:SuplierFrom;size:20" json:"supplier_from"`
	CustomerID     string     `gorm:"column:CustomerID;size:50" json:"customer_id"`
	Description    string     `gorm:"column:Description;size:100" json:"description"`
	DocumentNumber string     `gorm:"column:DocumentNumber;size:20" json:"document_number"`

	Details []OrderDetail `gorm:"foreignKey:OrderNo;references:OrderNo" json:"details,omitempty"`
}

func (OrderMaster) TableName() string { return "OrderMaster" }
