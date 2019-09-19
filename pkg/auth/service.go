package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// Service provides basic operations on chirp domain model.
type Service interface {
	GetStoredPasswordByUsername(username string) (string, error)
	ComparePassword(incomingPassword, storedPassword string) (bool, error)
}

type service struct {
	auth Repository
}

// NewService creates an instance of the service for user domain model with necessary dependencies.
func NewService(auth Repository) Service {
	return &service{
		auth: auth,
	}
}

func (s *service) GetStoredPasswordByUsername(username string) (string, error) {
	password, err := s.auth.GetPassword(username)
	if err != nil {
		return "", err
	}
	return password, nil
}

func (s *service) ComparePassword(incomingPassword, storedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(incomingPassword))
	if err != nil {
		return false, err
	}

	return true, nil
}
