package model

// Product maps to dbo.Product.
// Primary key is composite: ProductNumber + ProductID.
type Product struct {
	ProductNumber     int    `gorm:"column:ProductNumber;primaryKey" json:"product_number"`
	ProductID         string `gorm:"column:ProductID;primaryKey;size:25" json:"product_id"`
	ProductCategoryID *int   `gorm:"column:ProductCotegoryID" json:"product_category_id"`
	ProductName       string `gorm:"column:ProductName;size:50" json:"product_name"`
	Measure           string `gorm:"column:Measure;size:20" json:"measure"`
	Description       string `gorm:"column:Description;size:100" json:"description"`
	DocumentNumber    string `gorm:"column:DocumentNumber;size:20" json:"document_number"`

	Category *Category `gorm:"foreignKey:ProductCategoryID;references:CategoryID" json:"category,omitempty"`
}

func (Product) TableName() string { return "Product" }
