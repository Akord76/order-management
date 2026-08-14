package repository

import (
	"order-management/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create inserts a new AppUser via the usp_AppUser_CRUD stored procedure
// and returns the generated UserID.
func (r *UserRepository) Create(user *model.AppUser) (int, error) {
	var result struct {
		UserID int
	}
	err := r.db.Raw(
		`EXEC dbo.usp_AppUser_CRUD
			@Action = 'INSERT',
			@Username = ?,
			@PasswordHash = ?,
			@FullName = ?,
			@Email = ?,
			@RoleName = ?,
			@IsActive = ?`,
		user.Username, user.PasswordHash, user.FullName, user.Email, user.RoleName, user.IsActive,
	).Scan(&result).Error
	if err != nil {
		return 0, err
	}
	return result.UserID, nil
}

// FindByID retrieves a single AppUser (or the whole list if id is nil) via
// the usp_AppUser_CRUD stored procedure in GET mode.
func (r *UserRepository) FindByID(userID int) (*model.AppUser, error) {
	var user model.AppUser
	err := r.db.Raw(
		`EXEC dbo.usp_AppUser_CRUD @Action = 'GET', @UserID = ?`, userID,
	).Scan(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindAll() ([]model.AppUser, error) {
	var users []model.AppUser
	err := r.db.Raw(`EXEC dbo.usp_AppUser_CRUD @Action = 'GET'`).Scan(&users).Error
	return users, err
}

// FindByUsername uses the dedicated login stored procedure, which returns
// the password hash needed for credential verification.
func (r *UserRepository) FindByUsername(username string) (*model.AppUser, error) {
	var user model.AppUser
	err := r.db.Raw(`EXEC dbo.usp_AppUser_Login @Username = ?`, username).Scan(&user).Error
	if err != nil {
		return nil, err
	}
	if user.Username == "" {
		return nil, gorm.ErrRecordNotFound
	}
	return &user, nil
}

func (r *UserRepository) Update(user *model.AppUser) error {
	return r.db.Exec(
		`EXEC dbo.usp_AppUser_CRUD
			@Action = 'UPDATE',
			@UserID = ?,
			@FullName = ?,
			@Email = ?,
			@RoleName = ?,
			@IsActive = ?`,
		user.UserID, user.FullName, user.Email, user.RoleName, user.IsActive,
	).Error
}

func (r *UserRepository) Delete(userID int) error {
	return r.db.Exec(`EXEC dbo.usp_AppUser_CRUD @Action = 'DELETE', @UserID = ?`, userID).Error
}
