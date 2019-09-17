package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	"github.com/joshuabezaleel/chirp-server/pkg/core/user"
)

type userHandler struct {
	service user.Service
}

func (handler *userHandler) RegisterRouter(router *mux.Router) {
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

	err = handler.service.RegisterUser(request.Username, request.Email, handler.hashAndSalt([]byte(request.Password)), request.Role)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode("New user registered.")
}

func (handler *userHandler) hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}
