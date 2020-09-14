package handler

import (
	"dave-web-app/packages/session-service/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

type getSessionsBody struct {
	SessionIds []string `json:"sessionIds"`
}

func GetSessions(datastore service.Datastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var body getSessionsBody
		var sessions *[]service.Session
		err := json.NewDecoder(r.Body).Decode(&body)
		// if err == nil then the request body has speakerIds field
		// that is, get speakers by ids
		if err == nil {
			sessions, err = datastore.GetSessionsByIds(body.SessionIds)
		} else {
			// Reset the err var so that it doesn't trigger `Failed to get speakers` below
			err = nil
			sessions, err = datastore.GetAllSessions()
		}

		if err != nil {
			log.Printf("Failed to get sessions: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sessionBytes, err := json.Marshal(sessions)
		if err != nil {
			log.Printf("Failed to marshal sessions: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(sessionBytes)
	}
}
