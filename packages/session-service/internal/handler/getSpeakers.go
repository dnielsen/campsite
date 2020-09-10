package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetSpeakers(datastore SpeakerDatastore) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		speakers := datastore.GetAllSpeakers()
		speakerBytes, err := json.Marshal(speakers)
		if err != nil {
			log.Fatalf("Failed to marshal service: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(speakerBytes)
	}
}