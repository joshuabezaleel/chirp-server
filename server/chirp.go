package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"

	"github.com/joshuabezaleel/chirp-server/pkg/core/chirp"
)

type chirpHandler struct {
	service chirp.Service
}

func (handler *chirpHandler) registerRouter(router *mux.Router) {
	router.HandleFunc("/chirp", handler.newChirp).Methods("POST")
}

func (handler *chirpHandler) newChirp(w http.ResponseWriter, r *http.Request) {
	// Retrieve text.
	// text := r.FormValue("text")

	// Retrieve picture.
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return
	}
	// Delete cmd and add assets/
	dir = dir[:len(dir)-3]
	dir = dir + "assets/"

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileLocation := filepath.Join(dir, "images", header.Filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	_, err = io.Copy(targetFile, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("File uploaded")

	// Retrieve picture.
	// file, handle, err := r.FormFile("picture")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// defer file.Close()

	// fmt.Printf("Uploaded file: %+v\n", handle.Filename)
	// fmt.Printf("File size: %+v\n", handle.Size)
	// fmt.Printf("MIME header: %+v\n", handle.Header)

	// fileBytes, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	log.Println(err)
	// }

	// err = handler.service.Chirp(text, fileBytes)

	json.NewEncoder(w).Encode("Chirp chirped.")

}

// func (handler *userHandler) registerUser(w http.ResponseWriter, r *http.Request) {

// 	var request struct {
// 		Username string `json:"username"`
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 		Role     string `json:"role"`
// 	}

// 	err := json.NewDecoder(r.Body).Decode(&request)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	defer r.Body.Close()

// 	err = handler.service.RegisterUser(request.Username, request.Email, handler.hashAndSalt([]byte(request.Password)), request.Role)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	json.NewEncoder(w).Encode("New user registered.")
// }
