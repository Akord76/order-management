package model

import "time"

// Employee maps to dbo.Employee.
// Primary key is composite: EmployeeId (identity) + EmployeeCardNumber.
type Employee struct {
	EmployeeID         int        `gorm:"column:EmployeeId;primaryKey;autoIncrement" json:"employee_id"`
	EmployeeCardNumber int        `gorm:"column:EmployeeCardNumber;primaryKey" json:"employee_card_number"`
	FirstName          string     `gorm:"column:FirstName;size:25" json:"first_name"`
	LastName            string     `gorm:"column:LastName;size:25" json:"last_name"`
	DateOfBirth        *time.Time `gorm:"column:DateOfBrith" json:"date_of_birth"`
	Gender              string     `gorm:"column:Gender;size:25" json:"gender"`
	Email               string     `gorm:"column:Email;size:25" json:"email"`
	DepartmentID        int        `gorm:"column:DepartmentID;not null" json:"department_id"`
	PhoneNumber         string     `gorm:"column:PhoneNumber;size:50" json:"phone_number"`
	BankNumber          string     `gorm:"column:BankNumber;size:50" json:"bank_number"`
	PhotoPath           string     `gorm:"column:PhotoPath;size:50" json:"photo_path"`
}

func (Employee) TableName() string { return "Employee" }
