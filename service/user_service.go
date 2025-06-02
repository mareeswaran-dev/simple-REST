package service

import (
	"errors"
	"strings"

	"user-crud-api/models"
	"user-crud-api/repository"
)

// UserService handles business logic for users
type UserService struct {
	repo *repository.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser creates a new user after validating the input
func (s *UserService) CreateUser(user *models.User) error {
	if err := s.validateUser(user); err != nil {
		return err
	}
	return s.repo.Create(user)
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(id int) (*models.User, error) {
	return s.repo.GetByID(id)
}

// GetAllUsers retrieves all users
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

// UpdateUser updates an existing user after validating the input
func (s *UserService) UpdateUser(user *models.User) error {
	if err := s.validateUser(user); err != nil {
		return err
	}
	return s.repo.Update(user)
}

// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}

// validateUser performs validation on user data
func (s *UserService) validateUser(user *models.User) error {
	if strings.TrimSpace(user.Username) == "" {
		return errors.New("username is required")
	}
	if strings.TrimSpace(user.Email) == "" {
		return errors.New("email is required")
	}
	if !strings.Contains(user.Email, "@") {
		return errors.New("invalid email format")
	}
	if strings.TrimSpace(user.Password) == "" {
		return errors.New("password is required")
	}
	if len(user.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}
	return nil
}
