package errors

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound       = errors.New("not found")
	ErrorUnauthorized = errors.New("unauthorized")
	ErrInvalidInput   = errors.New("invalid input")
	ErrTimeout        = errors.New("operation timed out")
)

type UserService struct {
	Users map[int]string
}

func NewUserService(users map[int]string) *UserService {
	return &UserService{
		Users: users,
	}
}

func (s *UserService) GetUser(id int) (string, error) {
	if id < 0 {
		return "", fmt.Errorf("user id %d: %w", id, ErrInvalidInput)
	}

	user, exists := s.Users[id]

	if !exists {
		return "", fmt.Errorf("user id %d: %w", id, ErrNotFound)
	}

	return user, nil

}
func (s *UserService) DeleteUser(id int) (string, error) {
	if id < 0 {
		return "", fmt.Errorf("user id %d: %w", id, ErrInvalidInput)
	}

	user, exists := s.Users[id]

	if !exists {
		return "", fmt.Errorf("user id %d: %w", id, ErrNotFound)
	}

	return user, nil
}
