package user

import (
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// ErrDuplicate is used when a user with particular username/email already exists.
var ErrDuplicate = errors.New("User with username/email already exists")

// ErrWrongAuth is used when a user input a wrong combination of username/email and password.
var ErrWrongAuth = errors.New("Wrong combination of username/email and password")

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("Invalid argument")

// Service provides basic operations on user domain model.
type Service interface {
	RegisterUser(username, email, password, role string) error
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

// RegisterUser register a new user.
func (s *service) RegisterUser(username, email, password, role string) error {
	if len(username) == 0 {
		return ErrInvalidArgument
	}

	user := NewUser(username, email, s.hashAndSalt([]byte(password)), role, time.Now())

	return s.user.Create(user)
}

func (s *service) hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}
