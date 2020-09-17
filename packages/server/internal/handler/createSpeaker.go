
package handler

import (
	"dave-web-app/packages/server/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

func CreateSpeaker(datastore service.SpeakerDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i service.CreateSpeakerInput

		// Decode the body
		err := json.NewDecoder(r.Body).Decode(&i)
		if err != nil {
			log.Printf("Failed to unmarshal speaker input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Create the speaker in the database.
		speaker, err := datastore.CreateSpeaker(i)
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

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(speakerBytes)
	}
}