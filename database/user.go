package database

import (
	"database/sql"
	"time"

	user "github.com/joshuabezaleel/library-server/pkg/core/user"
)

type repository struct {
	DB *sql.DB
}

// NewRepository returns initialized implementations of the particular Repository.
func NewRepository(DB *sql.DB) user.Repository {
	return &repository{
		DB: DB,
	}
}

func (repo *repository) Create(user *user.User) error {
	_, err := repo.DB.Exec("INSERT INTO users (username, email, password, role, register_at) VALUES ($1, $2, $3, $4, $5)", user.Username, user.Email, user.Password, user.Role, time.Now())

	if err != nil {
		return err
	}

	return nil
}
