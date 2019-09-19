package chirp

import "time"

// Chirp defines a single chirp.
type Chirp struct {
	ChirpID uint64 `json:"chirpID"`
	UserID  int    `json:"userID"`
	Text    string `json:"text"`
	// MediaID   uint64    `json:"mediaID"`
	ChirpedAt time.Time `json:"chirpedAt"`
}

// NewChirp creates a new instance of chirp.
// func NewChirp(id uint64, text string, picture uint64) *Chirp {
// 	return &Chirp{
// 		ID:      id,
// 		Text:    text,
// 		Picture: picture,
// 	}
// }
func NewChirp(chirpID uint64, userID int, text string, chirpedAt time.Time) *Chirp {
	return &Chirp{
		ChirpID:   chirpID,
		UserID:    userID,
		Text:      text,
		ChirpedAt: chirpedAt,
	}
}
