package database

import (
	"database/sql"

	user "github.com/joshuabezaleel/chirp-server/pkg/core/user"
)

type userRepository struct {
	DB *sql.DB
}

// NewUserRepository returns initialized implementations of user repository.
func NewUserRepository(DB *sql.DB) user.Repository {
	return &userRepository{
		DB: DB,
	}
}

func (repo *userRepository) Create(user *user.User) error {
	_, err := repo.DB.Exec("INSERT INTO users (username, email, password, role, register_at) VALUES ($1, $2, $3, $4, $5)", user.Username, user.Email, user.Password, user.Role, user.RegisteredAt)

	if err != nil {
		return err
	}

	return nil
}
