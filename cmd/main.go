package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

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
	)
	userRepo = database.NewUserRepository(db)
	chirpRepo = database.NewChirpRepository(db)

	// Setting up domain services
	var (
		userService  user.Service
		chirpService chirp.Service
	)
	userService = user.NewService(userRepo)
	chirpService = chirp.NewService(chirpRepo)

	srv := server.New(userService, chirpService)

	err = http.ListenAndServe(port, srv.Router)
	if err != nil {
		log.Println(err)
	}
}
