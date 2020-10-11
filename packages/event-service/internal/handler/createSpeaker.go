package handler

import (
	"campsite/packages/event-service/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

// `/speakers` POST route. It communicates with the speaker service only.
func CreateSpeaker(api service.SpeakerAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the body.
		var i service.SpeakerInput
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			log.Printf("Failed to unmarshal speaker input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Request the speaker service to create a speaker.
		speaker, err := api.CreateSpeaker(i)
		if err != nil {
			log.Printf("Failed to create speaker: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Marshal the speaker.
		speakerBytes, err := json.Marshal(speaker)
		if err != nil {
			log.Printf("Failed to marshal speaker: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond JSON with the created speaker
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(speakerBytes)
	}
}
