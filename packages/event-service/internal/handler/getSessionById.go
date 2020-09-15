package handler

import (
	"dave-web-app/packages/event-service/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetSessionById(datastore service.Datastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]

		// Get the session from the session service.
		session, err := datastore.GetSessionById(id)
		if err != nil {
			log.Printf("Failed to get session: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Get the session speakers from speaker service since we're always displaying them along the full session
		// and attach them to the session.
		speakers, err := datastore.GetSpeakersByIds(session.SpeakerIds)
		if err != nil {
			log.Printf("Failed to get speakers: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		session.Speakers = *speakers

		// Marshal the session.
		sessionBytes, err := json.Marshal(session)
		if err != nil {
			log.Printf("Failed to marshal session: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond json with the session and the attached speakers.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(sessionBytes)
	}
}
