package handler

import (
	"dave-web-app/packages/speaker-service/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const ID = "id"

func GetSpeakerById(datastore service.SpeakerDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars[ID]
		speaker, err := datastore.GetSpeakerById(id)
		if err != nil {
			log.Printf("Failed to get speaker by id: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println(speaker.SessionIds)
		speakerBytes, err := json.Marshal(speaker)
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
