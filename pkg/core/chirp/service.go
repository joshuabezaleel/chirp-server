package chirp

import "errors"

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("Invalid argument")

// Service provides basic operations on chirp domain model.
type Service interface {
	Chirp(text string, pictures []string) error
}

type service struct {
	chirp Repository
}

// NewService creates an instance of the service for user domain model with necessary dependencies.
func NewService(chirp Repository) Service {
	return &service{
		chirp: chirp,
	}
}

// AddUser register a new user.
func (s *service) Chirp(text string, pictures []string) error {

	return nil
}
