package chirp

import "time"

// Chirp defines a single chirp.
type Chirp struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Text      string    `json:"text"`
	Pictures  []string  `json:"pictures"`
	ChirpedAt time.Time `json:"chirpedAt"`
}

// NewChirp creates a new instance of chirp.
func NewChirp(text string, pictures []string) *Chirp {
	return &Chirp{
		Text:     text,
		Pictures: pictures,
	}
}
