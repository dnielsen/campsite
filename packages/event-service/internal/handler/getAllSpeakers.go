package handler

import (
	"campsite/packages/event-service/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

// `/speakers` GET route. It communicates with the speaker service only.
func GetAllSpeakers(datastore service.SpeakerAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get all speakers from the speaker service.
		speakers, err := datastore.GetAllSpeakers()
		if err != nil {
			log.Printf("Failed to get speakers: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Marshal the speakers.
		speakerBytes, err := json.Marshal(speakers)
		if err != nil {
			log.Printf("Failed to marshal sessions: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond json with the speakers.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(speakerBytes)
	}
}
