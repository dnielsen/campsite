package handler

import (
	"campsite/packages/session-service/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

// `/` POST route.
func CreateSession(datastore service.SessionAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the body.
		var i service.SessionInput
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			log.Printf("Failed to unmarshal session input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Create the session in the database.
		session, err := datastore.CreateSession(i)
		if err != nil {
			log.Printf("Failed to create session: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Marshal the created session.
		sessionBytes, err := json.Marshal(session)
		if err != nil {
			log.Printf("Failed to marshal session: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond JSON with the created session.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(sessionBytes)
	}
}