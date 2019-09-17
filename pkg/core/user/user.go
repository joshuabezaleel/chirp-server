package user

// User defines a user.
type User struct {
	// ID         int       `json:"id"`
	Username string `json:"username"`
	// Email      string    `json:"email"`
	// Password   string    `json:"password"`
	// Role       string    `json:"role"`
	// RegisterAt time.Time `json:"registerAt"`
}

// NewUser creates a new instance of user.
func NewUser(username string) *User {
	return &User{
		Username: username,
	}
}

// func NewUser(id int, username, email, password, role string, registerAt time.Time) *User {
// 	return &User{
// 		ID:         id,
// 		Username:   username,
// 		Email:      email,
// 		Password:   password,
// 		Role:       role,
// 		RegisterAt: registerAt,
// 	}
// }

// // GetRole return role of a user.
// func (u *User) GetRole() string {
// 	return u.Role
// }
