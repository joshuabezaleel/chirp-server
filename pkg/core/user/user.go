package user

import "time"

// User defines a user.
type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Role         string    `json:"role"`
	RegisteredAt time.Time `json:"registeredAt"`
}

// NewUser creates a new instance of user.
func NewUser(username, email, password, role string, registeredAt time.Time) *User {
	return &User{
		Username:     username,
		Email:        email,
		Password:     password,
		Role:         role,
		RegisteredAt: registeredAt,
	}
}

// // GetRole return role of a user.
// func (u *User) GetRole() string {
// 	return u.Role
// }
