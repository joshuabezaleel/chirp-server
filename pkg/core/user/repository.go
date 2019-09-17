package user

// Repository provides access to user store.
type Repository interface {
	Create(user *User) error
	// FindByID(id int) (*User, error)
	// FindByEmail(email string) (*User, error)
	// FindByUsername(username string) (*User, error)
	// GetRoleFromUsername(username string) (role string)
}
