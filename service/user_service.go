package service

import (
	"errors"

	"order-management/auth"
	"order-management/model"
	"order-management/repository"
)

// UserService backs the ADMIN "User Management" capability:
// Create User, Update User, Delete User, View Users.
// (Login/self-registration lives in AuthService instead.)
type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Create is used by an ADMIN to provision a new user account, optionally
// assigning any role (ADMIN, MANAGER, USER).
func (s *UserService) Create(user *model.AppUser, plainPassword string) (int, error) {
	if user.Username == "" || plainPassword == "" {
		return 0, errors.New("username and password are required")
	}
	if user.RoleName == "" {
		user.RoleName = RoleUser
	}
	if !isValidRole(user.RoleName) {
		return 0, errors.New("role must be one of ADMIN, MANAGER, USER")
	}

	hash, err := auth.HashPassword(plainPassword)
	if err != nil {
		return 0, err
	}
	user.PasswordHash = hash

	return s.repo.Create(user)
}

func (s *UserService) GetAll() ([]model.AppUser, error) {
	return s.repo.FindAll()
}

func (s *UserService) GetByID(userID int) (*model.AppUser, error) {
	return s.repo.FindByID(userID)
}

func (s *UserService) Update(user *model.AppUser) error {
	if user.UserID == 0 {
		return errors.New("user id is required")
	}
	if user.RoleName != "" && !isValidRole(user.RoleName) {
		return errors.New("role must be one of ADMIN, MANAGER, USER")
	}
	return s.repo.Update(user)
}

func (s *UserService) Delete(userID int) error {
	return s.repo.Delete(userID)
}

const (
	RoleAdmin   = "ADMIN"
	RoleManager = "MANAGER"
	RoleUser    = "USER"
)

func isValidRole(role string) bool {
	switch role {
	case RoleAdmin, RoleManager, RoleUser:
		return true
	default:
		return false
	}
}
