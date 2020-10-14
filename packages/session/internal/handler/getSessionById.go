package handler

import (
	"campsite/packages/session/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetSessionById(datastore service.SessionAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]

		// Get the session from the database.
		session, err := datastore.GetSessionById(id)
		if err != nil {
			log.Printf("Failed to get session: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Marshal the session.
		sessionBytes, err := json.Marshal(session)
		if err != nil {
			log.Printf("Failed to marshal session: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond JSON with the session.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(sessionBytes)
	}
}
