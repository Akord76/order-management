package model

// Category maps to dbo.Category.
type Category struct {
	CategoryID   int    `gorm:"column:CategoryId;primaryKey;autoIncrement" json:"category_id"`
	CategoryName string `gorm:"column:CategoryName;size:25" json:"category_name"`
}

func (Category) TableName() string { return "Category" }
