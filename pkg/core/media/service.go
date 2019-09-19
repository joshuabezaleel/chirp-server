package media

import (
	"errors"
	"math/rand"
	"time"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("Invalid argument")

// Service provides basic operations on media domain model.
type Service interface {
	UploadMedia(size int) error
}

type service struct {
	media Repository
}

// NewService creates an instance of the service for media domain model with necessary dependencies.
func NewService(media Repository) Service {
	return &service{
		media: media,
	}
}

// UploadMedia is used for uploading a new media.
func (s *service) UploadMedia(size int) error {
	media := NewMedia(s.generateMediaID(), size, time.Now())

	return s.media.Create(media)
}

// generateMediaID used to generate random uint64 for the media ID.
func (s *service) generateMediaID() uint64 {
	return rand.Uint64()
}
