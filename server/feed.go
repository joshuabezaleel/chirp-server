package server

import (
	"encoding/json"
	"net/http"

	"github.com/joshuabezaleel/chirp-server/pkg/auth"

	"github.com/gorilla/mux"

	"github.com/joshuabezaleel/chirp-server/pkg/feed"
)

type feedHandler struct {
	service feed.Service
	auth    auth.Service
}

func (handler *feedHandler) registerRouter(router *mux.Router) {
	router.HandleFunc("/index", handler.auth.AuthenticationMiddleware(handler.feedIndex)).Methods("POST")
}

func (handler *feedHandler) feedIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Feed index")
}
