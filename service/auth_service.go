package service

import (
	"errors"

	"order-management/auth"
	"order-management/model"
	"order-management/repository"
)

type AuthService struct {
	repo        *repository.UserRepository
	jwtSecret   string
	jwtExpireHr int
}

func NewAuthService(repo *repository.UserRepository, jwtSecret string, jwtExpireHr int) *AuthService {
	return &AuthService{repo: repo, jwtSecret: jwtSecret, jwtExpireHr: jwtExpireHr}
}

// Register creates a new AppUser with a securely hashed password.
func (s *AuthService) Register(user *model.AppUser, plainPassword string) (int, error) {
	if user.Username == "" || plainPassword == "" {
		return 0, errors.New("username and password are required")
	}

	hash, err := auth.HashPassword(plainPassword)
	if err != nil {
		return 0, err
	}
	user.PasswordHash = hash

	if user.RoleName == "" {
		user.RoleName = "USER"
	}

	return s.repo.Create(user)
}

// Login validates credentials and returns a signed JWT on success.
func (s *AuthService) Login(username, plainPassword string) (string, *model.AppUser, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", nil, errors.New("invalid username or password")
	}

	if !user.IsActive {
		return "", nil, errors.New("user account is inactive")
	}

	if !auth.CheckPassword(user.PasswordHash, plainPassword) {
		return "", nil, errors.New("invalid username or password")
	}

	token, err := auth.GenerateToken(s.jwtSecret, s.jwtExpireHr, user.UserID, user.Username, user.RoleName)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}
