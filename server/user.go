package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/joshuabezaleel/library-server/pkg/core/user"
)

type userHandler struct {
	service user.Service
}

// func (handler *userHandler) router() *mux.Router {
// 	root := mux.NewRouter()
// 	router := root.PathPrefix("/user").Subrouter()

// 	router.HandleFunc("/register", handler.registerUser).Methods("POST")

// 	return root
// }

func (handler *userHandler) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/register", handler.registerUser).Methods("POST")
}

func (handler *userHandler) registerUser(w http.ResponseWriter, r *http.Request) {
	// ctx := context.Background()

	var request struct {
		Username string
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer r.Body.Close()

	err = handler.service.RegisterUser(request.Username)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode("New user registered.")
}
