package database

import (
	"database/sql"

	chirp "github.com/joshuabezaleel/chirp-server/pkg/core/chirp"
)

type chirpRepository struct {
	DB *sql.DB
}

// NewChirpRepository returns initialized implementations of chirp repository.
func NewChirpRepository(DB *sql.DB) chirp.Repository {
	return &chirpRepository{
		DB: DB,
	}
}

// func (repo *repository) Create(user *user.User) error {
// 	_, err := repo.DB.Exec("INSERT INTO users (username, email, password, role, register_at) VALUES ($1, $2, $3, $4, $5)", user.Username, user.Email, user.Password, user.Role, time.Now())

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (repo *chirpRepository) Create(chirp *chirp.Chirp) error {
	return nil
}
