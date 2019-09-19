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

func (repo *authRepository) GetPasswordByUsername(username string) (string, error) {
	var password string
	// fmt.Println(password)

	res := repo.DB.QueryRow("SELECT password FROM users WHERE username=$1", username)
	err := res.Scan(&password)
	if err != nil {
		return "", err
	}
	// fmt.Println(password)

	return password, nil
}
