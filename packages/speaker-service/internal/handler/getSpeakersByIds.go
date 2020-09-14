package handler

import (
	"dave-web-app/packages/speaker-service/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

type getSpeakersBody struct {
	SpeakerIds []string `json:"speakerIds"`
}

func GetSpeakersByIds(datastore service.SpeakerDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body getSpeakersBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			log.Printf("Failed to decode request body: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		log.Println(body.SpeakerIds)


		speakers, err := datastore.GetSpeakersByIds(body.SpeakerIds)
		if err != nil {
			log.Printf("Failed to get speakers with specified ids: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
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
