package model

// Company maps to dbo.Company.
// Primary key is composite: CompanyNumber + CompanyID.
type Company struct {
	CompanyNumber int    `gorm:"column:CompanyNumber;primaryKey" json:"company_number"`
	CompanyID     string `gorm:"column:CompanyID;primaryKey;size:50" json:"company_id"`
	CompanyName   string `gorm:"column:CompanyName;size:50" json:"company_name"`
	Address       string `gorm:"column:Address;size:100" json:"address"`
	ContactPerson string `gorm:"column:ContactPerson;size:50" json:"contact_person"`
}

func (Company) TableName() string { return "Company" }
