package service

import (
	"auth-service/internal/model"
	"auth-service/internal/repository"
	"time"
)

// AuthService coordinates business logic related to authentication and
// user management. It depends on a repository for persistence.
type AuthService struct {
	repo *repository.UserRepository
}

// NewAuthService constructs an AuthService with the provided repository.
func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

// Register creates a new user record. Password hashing can be added later.
func (s *AuthService) Register(email string, password string) (*model.User, error) {
	user := &model.User{
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
