package handler

import (
	"dave-web-app/packages/speaker-service/internal/service"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const ID = "id"

func GetSpeakersByEventId(datastore service.SpeakerDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		strId := vars[ID]
		id, err := uuid.Parse(strId)
		if err != nil {
			log.Printf("Failed to convert id to uuid")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		speakers, err := datastore.GetSpeakersByIds(id)
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
