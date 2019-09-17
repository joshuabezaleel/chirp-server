package user

import (
	"errors"
)

// ErrDuplicate is used when a user already exists.
var ErrDuplicate = errors.New("User already exists")

// ErrWrongAuth is used when a user input a wrong combination of username/email and password.
var ErrWrongAuth = errors.New("Wrong combination of username/email and password")

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("Invalid argument")

// Service provides basic operations on user domain model.
type Service interface {
	RegisterUser(username string) error
}

type service struct {
	user Repository
}

// NewService creates an instance of the service for user domain model with necessary dependencies.
func NewService(user Repository) Service {
	return &service{
		user: user,
	}
}

// AddUser register a new user.
func (s *service) RegisterUser(username string) error {
	if len(username) == 0 {
		return ErrInvalidArgument
	}

	user := NewUser(username)

	return s.user.Create(user)
}
