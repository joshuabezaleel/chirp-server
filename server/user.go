package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/joshuabezaleel/chirp-server/pkg/core/user"
)

type userHandler struct {
	service user.Service
}

func (handler *userHandler) registerRouter(router *mux.Router) {
	router.HandleFunc("/register", handler.registerUser).Methods("POST")
}

func (handler *userHandler) registerUser(w http.ResponseWriter, r *http.Request) {

	var request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		return
	}
	defer r.Body.Close()

	err = handler.service.RegisterUser(request.Username, request.Email, request.Password, request.Role)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode("New user registered.")
}
