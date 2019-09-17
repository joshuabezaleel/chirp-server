package database

import (
	"database/sql"

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
	_, err := repo.DB.Exec("INSERT INTO users (username) VALUES ($1)", user.Username)

	if err != nil {
		return err
	}

	return nil
}
