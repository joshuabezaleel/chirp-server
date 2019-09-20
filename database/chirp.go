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

func (repo *chirpRepository) Create(chirp *chirp.Chirp) error {
	return nil
}
