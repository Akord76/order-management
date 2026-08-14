package model

import "time"

// AppUser maps to dbo.AppUser (referenced by usp_AppUser_CRUD / usp_AppUser_Login).
type AppUser struct {
	UserID       int       `gorm:"column:UserID;primaryKey;autoIncrement" json:"user_id"`
	Username     string    `gorm:"column:Username;size:50;uniqueIndex" json:"username"`
	PasswordHash string    `gorm:"column:PasswordHash;size:255" json:"-"`
	FullName     string    `gorm:"column:FullName;size:100" json:"full_name"`
	Email        string    `gorm:"column:Email;size:100" json:"email"`
	RoleName     string    `gorm:"column:RoleName;size:30" json:"role_name"`
	IsActive     bool      `gorm:"column:IsActive" json:"is_active"`
	CreatedAt    time.Time `gorm:"column:CreatedAt;autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:UpdatedAt;autoUpdateTime" json:"updated_at"`
}

func (AppUser) TableName() string { return "AppUser" }
