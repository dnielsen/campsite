package handler

import (
	"dave-web-app/packages/event-service/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const ID = "id"


func GetEventById(datastore service.Datastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars[ID]
		event, err := datastore.GetEventById(id)
		if err != nil {
			log.Printf("Failed to get event: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sessions, err := datastore.GetSessionsByIds(event.SessionIds)
		if err != nil {
			log.Printf("Failed to get sessions: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		event.Sessions = *sessions

		speakers, err := datastore.GetSpeakersByIds(event.SpeakerIds)
		if err != nil {
			log.Printf("Failed to get speakers: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		event.Speakers = *speakers

		eventBytes, err := json.Marshal(event)
		if err != nil {
			log.Printf("Failed to marshal full event: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(eventBytes)
	}
}
