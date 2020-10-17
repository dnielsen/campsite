package handler

import (
	"campsite/packages/speaker/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// `/{id}` GET route.
func GetSpeakerById(datastore service.SpeakerAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the speaker id parameter
		vars := mux.Vars(r)
		id := vars[ID]

		// Get the speaker from the database.
		speaker, err := datastore.GetSpeakerById(id)
		if err != nil {
			log.Printf("Failed to get speaker by id: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Marshal the speaker.
		speakerBytes, err := json.Marshal(speaker)
		if err != nil {
			log.Printf("Failed to marshal speakers: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond json with the speaker.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(speakerBytes)
	}
}
