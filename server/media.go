package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joshuabezaleel/chirp-server/pkg/core/media"
)

type mediaHandler struct {
	service media.Service
}

func (handler *mediaHandler) registerRouter(router *mux.Router) {
	router.HandleFunc("/upload", handler.uploadMedia).Methods("POST")
}

func (handler *mediaHandler) uploadMedia(w http.ResponseWriter, r *http.Request) {

	// var request struct {
	// 	Username string `json:"username"`
	// 	Email    string `json:"email"`
	// 	Password string `json:"password"`
	// 	Role     string `json:"role"`
	// }

	// err := json.NewDecoder(r.Body).Decode(&request)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// defer r.Body.Close()

	// err = handler.service.RegisterUser(request.Username, request.Email, request.Password, request.Role)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// json.NewEncoder(w).Encode("New user registered.")
}
