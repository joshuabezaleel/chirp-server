package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/joshuabezaleel/chirp-server/pkg/feed"

	"github.com/joshuabezaleel/chirp-server/pkg/auth"

	_ "github.com/lib/pq"

	"github.com/joshuabezaleel/chirp-server/database"
	"github.com/joshuabezaleel/chirp-server/pkg/core/chirp"
	"github.com/joshuabezaleel/chirp-server/pkg/core/user"
	"github.com/joshuabezaleel/chirp-server/server"
)

const (
	port         = ":8082"
	connhost     = "localhost"
	connport     = 8081
	connusername = "postgres"
	connpassword = "postgres"
	dbname       = "chirp-server-18-9"
)

func main() {
	connectionString := fmt.Sprintf("host = %s port=%d user=%s password=%s dbname=%s sslmode=disable", connhost, connport, connusername, connpassword, dbname)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	// Setting up domain repositories
	var (
		userRepo  user.Repository
		chirpRepo chirp.Repository
		authRepo  auth.Repository
		feedRepo  feed.Repository
	)
	userRepo = database.NewUserRepository(db)
	chirpRepo = database.NewChirpRepository(db)
	authRepo = database.NewAuthRepository(db)
	feedRepo = database.NewFeedRepository(db)

	// Setting up domain services
	var (
		userService  user.Service
		chirpService chirp.Service
		authService  auth.Service
		feedService  feed.Service
	)
	userService = user.NewService(userRepo)
	chirpService = chirp.NewService(chirpRepo)
	authService = auth.NewService(authRepo)
	feedService = feed.NewService(feedRepo)

	srv := server.New(userService, chirpService, authService, feedService)

	err = http.ListenAndServe(port, srv.Router)
	if err != nil {
		log.Println(err)
	}
}
