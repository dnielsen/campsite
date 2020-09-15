package handler

import (
	"dave-web-app/packages/session-service/internal/service"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// the `/` route. You could specify an id parameter like so: `/?id=453,123,435` which would fetch the
// speakers with the ids of 453, 123, and 435.
func GetSessions(datastore service.Datastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the ids from the query parameter.
		idsString := r.URL.Query().Get(ID)
		ids := strings.Split(idsString, ",")

		// Get the sessions from the database.
		var err error
		var sessions *[]service.Session
		// If there are no parameters, by default there's a string array with an empty string inside.
		// So, we're checking for an array of length 1 with an empty string inside.
		idsNotSpecified := len(ids) == 1 && len(ids[0]) == 0
		if idsNotSpecified {
			sessions, err = datastore.GetAllSessions()
		} else {
			sessions, err = datastore.GetSessionsByIds(ids)
		}

		if err != nil {
			log.Printf("Failed to get sessions: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Marshal the sessions
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
