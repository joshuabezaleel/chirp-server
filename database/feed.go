package database

import (
	"database/sql"

	"github.com/joshuabezaleel/chirp-server/pkg/feed"
)

type feedRepository struct {
	DB *sql.DB
}

// NewFeedRepository returns initialized implementations of feed repository.
func NewFeedRepository(DB *sql.DB) feed.Repository {
	return &feedRepository{
		DB: DB,
	}
}
