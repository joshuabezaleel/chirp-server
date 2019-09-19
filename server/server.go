package server

import (
	"github.com/gorilla/mux"
	"github.com/joshuabezaleel/chirp-server/pkg/core/chirp"
	"github.com/joshuabezaleel/chirp-server/pkg/core/user"
)

// Server holds dependencies for a HTTP server.
type Server struct {
	User  user.Service
	Chirp chirp.Service

	Router *mux.Router
}

// New returns a new HTTP server.
func New(user user.Service, chirp chirp.Service) *Server {
	server := &Server{
		User:  user,
		Chirp: chirp,
	}

	router := mux.NewRouter()
	uh := userHandler{user}
	ch := chirpHandler{chirp}

	uh.registerRouter(router.PathPrefix("/users").Subrouter())
	ch.registerRouter(router.PathPrefix("/chirps").Subrouter())

	server.Router = router

	return server
}
