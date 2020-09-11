package handler

import (
	"dave-web-app/packages/speaker-service/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

func GetSpeakers(datastore service.SpeakerDatastore) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		speakers, err := datastore.GetAllSpeakers()
		if err != nil {
			log.Printf("Failed to get all speakers: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		speakerBytes, err := json.Marshal(speakers)
		if err != nil {
			log.Printf("Failed to marshal speakers: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(speakerBytes)
	}
}