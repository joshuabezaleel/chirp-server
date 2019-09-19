package chirp

import (
	"errors"
	"math/rand"
	"time"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("Invalid argument")

// Service provides basic operations on chirp domain model.
type Service interface {
	// Chirp(id uint64, text string, picture uint64) error
	Chirp(userID int, text string) error
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

// Chirp add a new chirp.
// func (s *service) Chirp(id uint64, text string, picture []byte) error {
// 	return s.chirp.Create(NewChirp(id, text, picture))
// }
func (s *service) Chirp(userID int, text string) error {
	chirp := NewChirp(s.generateChirpID(), userID, text, time.Now())

	return s.chirp.Create(chirp)
}

// generateChirpID used to generate random uint64 for the chirp ID.
func (s *service) generateChirpID() uint64 {
	return rand.Uint64()
}
