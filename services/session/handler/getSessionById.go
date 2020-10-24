package handler

import (
	"encoding/json"
	"github.com/dnielsen/campsite/services/session/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetSessionById(api service.SessionAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]
		// Get the session from the database.
		session, err := api.GetSessionById(id)
		if err != nil {
			log.Printf("Failed to get session: %v", err)
			http.NotFound(w, r)
			return
		}
		// Marshal the session.
		b, err := json.Marshal(session)
		if err != nil {
			log.Printf("Failed to marshal session: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond JSON with the session.
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
