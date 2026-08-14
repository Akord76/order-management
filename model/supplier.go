package model

// Supplier maps to dbo.Suplier.
// Primary key is composite: SuplierNumber (identity) + SuplierID.
type Supplier struct {
	SupplierNumber int    `gorm:"column:SuplierNumber;primaryKey;autoIncrement" json:"supplier_number"`
	SupplierID     string `gorm:"column:SuplierID;primaryKey;size:20" json:"supplier_id"`
	SupplierName   string `gorm:"column:SuplierName;size:50" json:"supplier_name"`
	Address        string `gorm:"column:Address;size:100" json:"address"`
	City           string `gorm:"column:City;size:25" json:"city"`
	ContactPerson  string `gorm:"column:ContactPerson;size:25" json:"contact_person"`
	Email          string `gorm:"column:Email;size:50" json:"email"`
	PhoneNumber    string `gorm:"column:PhoneNumber;size:20" json:"phone_number"`
	BankNumber     string `gorm:"column:BankNumber;size:20" json:"bank_number"`
}

func (Supplier) TableName() string { return "Suplier" }
