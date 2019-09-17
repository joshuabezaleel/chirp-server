package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/joshuabezaleel/library-server/database"
	"github.com/joshuabezaleel/library-server/pkg/core/user"
	"github.com/joshuabezaleel/library-server/server"
)

const (
	port         = ":8082"
	connhost     = "localhost"
	connport     = 8081
	connusername = "postgres"
	connpassword = "postgres"
	dbname       = "library-17-9"
)

func main() {
	connectionString := fmt.Sprintf("host = %s port=%d user=%s password=%s dbname=%s sslmode=disable", connhost, connport, connusername, connpassword, dbname)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Setting up domain repositories
	var (
		userRepo user.Repository
	)
	userRepo = database.NewRepository(db)

	// Setting up domain services
	var (
		userService user.Service
	)
	userService = user.NewService(userRepo)

	srv := server.New(userService)

	err = http.ListenAndServe(port, srv.Router)
	if err != nil {
		log.Fatal(err)
	}
}
