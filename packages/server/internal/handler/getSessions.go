package handler

import (
	"dave-web-app/packages/server/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

// `/sessions` GET route.
func GetSessions(api service.SessionAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// We can later add url parameters to
		// specify the ids of the events we want
		// if needed.

		// Get the sessions from the database.
		sessions, err := api.GetAllSessions()
		if err != nil {
			log.Printf("Failed to get sessions: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Marshal the sessions.
		sessionBytes, err := json.Marshal(sessions)
		if err != nil {
			log.Printf("Failed to marshal session: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond JSON with the sessions
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(sessionBytes)
	}
}
