package media

import "time"

// Media defines a media.
type Media struct {
	MediaID   uint64    `json:"mediaID"`
	Size      int       `json:"size"`
	CreatedAt time.Time `json:"createdAt"`
}

// NewMedia creates a new instance of media.
func NewMedia(mediaID uint64, size int, createdAt time.Time) *Media {
	return &Media{
		MediaID:   mediaID,
		Size:      size,
		CreatedAt: createdAt,
	}
}
