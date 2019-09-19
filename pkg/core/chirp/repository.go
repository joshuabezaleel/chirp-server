package chirp

// Repository provides access to the chirp store.
type Repository interface {
	Create(chirp *Chirp) error
	// Create(c) error
}
