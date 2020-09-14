package handler

import (
	"dave-web-app/packages/speaker-service/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

func GetSpeakers(datastore service.SpeakerDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var body getSpeakersBody
		var speakers *[]service.Speaker
		err := json.NewDecoder(r.Body).Decode(&body)
		// if err == nil then the request body has speakerIds field
		// that is, get speakers by ids
		if err == nil {
			speakers, err = datastore.GetSpeakersByIds(body.SpeakerIds)
		} else {
			// Reset the err var so that it doesn't trigger `Failed to get speakers` below
			err = nil
			speakers, err = datastore.GetAllSpeakers()
		}

		if err != nil {
			log.Printf("Failed to get speakers: %v", err)
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
