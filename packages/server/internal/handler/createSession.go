
package handler

import (
	"dave-web-app/packages/server/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

func CreateSession(datastore service.SessionDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i service.CreateSessionInput

		// Decode the body
		err := json.NewDecoder(r.Body).Decode(&i)
		if err != nil {
			log.Printf("Failed to unmarshal session input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Create the session in the database.
		session, err := datastore.CreateSession(i)
		if err != nil {
			log.Printf("Failed to create event: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sessionBytes, err := json.Marshal(session)
		if err != nil {
			log.Printf("Failed to marshal event: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(sessionBytes)
	}
}