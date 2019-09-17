package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joshuabezaleel/library-server/pkg/core/user"
)

// Server holds dependencies for a HTTP server.
type Server struct {
	User user.Service

	Router *mux.Router
}

// New returns a new HTTP server.
func New(user user.Service) *Server {
	server := &Server{
		User: user,
	}

	router := mux.NewRouter()
	uh := userHandler{user}

	router.HandleFunc("/", testHandler)
	uh.RegisterRouter(router.PathPrefix("/user").Subrouter())

	server.Router = router

	return server
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	return
}
