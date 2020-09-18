package handler

import (
	"dave-web-app/packages/server/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

func GetSpeakers(datastore service.SpeakerDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// We can later add url parameters to
		// specify the ids of the events we want
		// if needed.

		// Get the speakers from the database.
		speakers, err := datastore.GetAllSpeakers()
		if err != nil {
			log.Printf("Failed to get sessions: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Marshal the sessions.
		speakerBytes, err := json.Marshal(speakers)
		if err != nil {
			log.Printf("Failed to marshal speakers: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond JSON with the sessions
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(speakerBytes)
	}
}
