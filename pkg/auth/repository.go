package auth

// Repository provides access to the store by authentication service.
type Repository interface {
	GetPassword(username string) (string, error)
	// Create() error
	// Create(c) error
}
