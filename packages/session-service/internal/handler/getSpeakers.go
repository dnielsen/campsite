package handler

import (
	"../service"
	"encoding/json"
	"log"
	"net/http"
)

func GetSpeakers(datastore service.SpeakerDatastore) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		speakers, err := datastore.GetAllSpeakers()
		if err != nil {
			log.Fatalf("Failed to get all speakers: %v", err)
		}
		speakerBytes, err := json.Marshal(speakers)
		if err != nil {
			log.Fatalf("Failed to marshal speakers: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(speakerBytes)
	}
}