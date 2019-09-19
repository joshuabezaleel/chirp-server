package server

import (
	"encoding/json"
	"net/http"

	"github.com/joshuabezaleel/chirp-server/pkg/auth"

	"github.com/gorilla/mux"
	"github.com/joshuabezaleel/chirp-server/pkg/core/chirp"
	"github.com/joshuabezaleel/chirp-server/pkg/core/user"
)

// Server holds dependencies for a HTTP server.
type Server struct {
	User  user.Service
	Chirp chirp.Service
	Auth  auth.Service

	Router *mux.Router
}

// New returns a new HTTP server.
func New(user user.Service, chirp chirp.Service, auth auth.Service) *Server {
	server := &Server{
		User:  user,
		Chirp: chirp,
		Auth:  auth,
	}

	router := mux.NewRouter()
	uh := userHandler{user}
	ch := chirpHandler{chirp}
	ah := authHandler{auth}

	uh.registerRouter(router.PathPrefix("/users").Subrouter())
	ch.registerRouter(router.PathPrefix("/chirps").Subrouter())
	ah.registerRouter(router.PathPrefix("/auth").Subrouter())

	server.Router = router

	return server
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"Error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
