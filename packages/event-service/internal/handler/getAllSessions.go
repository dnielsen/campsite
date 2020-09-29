package handler

import (
	"campsite/packages/event-service/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

func GetAllSessions(datastore service.SessionAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get all sessions from the session service.
		sessions, err := datastore.GetAllSessions()
		if err != nil {
			log.Printf("Failed to get sessions: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Marshal the sessions.
		sessionBytes, err := json.Marshal(sessions)
		if err != nil {
			log.Printf("Failed to marshal sessions: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond json with the sessions.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(sessionBytes)
	}
}
