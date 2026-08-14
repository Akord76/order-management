package model

// Customer maps to dbo.Customer.
// Primary key is composite: CustNumber + CustID.
type Customer struct {
	CustNumber    int    `gorm:"column:CustNumber;primaryKey" json:"cust_number"`
	CustID        string `gorm:"column:CustID;primaryKey;size:50" json:"cust_id"`
	CustName      string `gorm:"column:CustName;size:50" json:"cust_name"`
	Address       string `gorm:"column:Address;size:255" json:"address"`
	City          string `gorm:"column:City;size:25" json:"city"`
	ContactPerson string `gorm:"column:ContactPerson;size:25" json:"contact_person"`
	Email         string `gorm:"column:Email;size:25" json:"email"`
	BankNumber    string `gorm:"column:BankNumber;size:50" json:"bank_number"`
	PhoneNumber   string `gorm:"column:PhoneNumber;size:50" json:"phone_number"`
}

func (Customer) TableName() string { return "Customer" }
