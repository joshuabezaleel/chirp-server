package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joshuabezaleel/chirp-server/pkg/auth"
)

type authHandler struct {
	service auth.Service
}

func (handler *authHandler) registerRouter(router *mux.Router) {
	router.HandleFunc("/login", handler.login).Methods("POST")
}

func (handler *authHandler) login(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	storedPassword, err := handler.service.GetStoredPasswordByUsername(request.Username)
	if err != nil {
		log.Println(err)
	}

	_, err = handler.service.ComparePassword(request.Password, storedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Password is wrong")
		return
	} else {
		fmt.Println("Same password!")
	}

	fmt.Println(storedPassword)
}
