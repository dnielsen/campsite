package handler

import (
	"encoding/json"
	"github.com/dnielsen/campsite/services/event/service"
	"log"
	"net/http"
)

// `/sessions` GET route. It communicates with the session service only.
func GetAllSessions(api service.SessionAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get all sessions from the session service.
		sessions, err := api.GetAllSessions()
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
