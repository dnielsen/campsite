package handler

import (
	"dave-web-app/packages/server/internal/service"
	"encoding/json"
	"github.com/go-playground/validator"
	"log"
	"net/http"
)

// `/speakers` POST route.
func CreateSpeaker(api service.SpeakerAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the body.
		var i service.SpeakerInput
		err := json.NewDecoder(r.Body).Decode(&i)
		if err != nil {
			log.Printf("Failed to unmarshal speaker input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Validate the input.
		if err := validator.New().Struct(i); err != nil {
			log.Printf("Failed to validate speaker input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Create the speaker in the database.
		speaker, err := api.CreateSpeaker(i)
		if err != nil {
			log.Printf("Failed to create speaker: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		speakerBytes, err := json.Marshal(speaker)
		if err != nil {
			log.Printf("Failed to marshal speaker: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond JSON with the created speaker.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(speakerBytes)
	}
}
