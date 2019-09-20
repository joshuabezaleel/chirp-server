package database

import (
	"database/sql"

	"github.com/joshuabezaleel/chirp-server/pkg/auth"
)

type authRepository struct {
	DB *sql.DB
}

// NewAuthRepository returns authentication service that provides access to the database.
func NewAuthRepository(DB *sql.DB) auth.Repository {
	return &authRepository{
		DB: DB,
	}
}

func (repo *authRepository) GetPassword(username string) (string, error) {
	var password string

	res := repo.DB.QueryRow("SELECT password FROM users WHERE username=$1", username)
	err := res.Scan(&password)
	if err != nil {
		return "", err
	}

	return password, nil
}
