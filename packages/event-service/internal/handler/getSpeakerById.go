package handler

import (
	"dave-web-app/packages/event-service/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const SPEAKER_ID = "speakerId"

func GetSpeakerById(datastore service.Datastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars[SPEAKER_ID]
		speaker, err := datastore.GetSpeakerById(id)
		if err != nil {
			log.Printf("Failed to get speaker: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		speakerBytes, err := json.Marshal(speaker)
		if err != nil {
			log.Printf("Failed to marshal speaker: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(speakerBytes)
	}
}
