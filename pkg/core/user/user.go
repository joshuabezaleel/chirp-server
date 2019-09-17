package user

import "time"

// User defines a user.
type User struct {
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Role       string    `json:"role"`
	RegisterAt time.Time `json:"registerAt"`
}

// NewUser creates a new instance of user.
func NewUser(username, email, password, role string) *User {
	return &User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
	}
}

// // GetRole return role of a user.
// func (u *User) GetRole() string {
// 	return u.Role
// }
