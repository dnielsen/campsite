package handler

import (
	"dave-web-app/packages/event-service/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetEventById(datastore service.Datastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("hlelo world")
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]

		// Get the event from the database.
		event, err := datastore.GetEventById(id)
		if err != nil {
			log.Printf("Failed to get event: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Get the sessions from the session service and add them to the event.
		sessions, err := datastore.GetSessionsByIds(event.SessionIds)
		if err != nil {
			log.Printf("Failed to get sessions: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		event.Sessions = *sessions

		// Get the speakers from the speaker service and add them to the event.
		speakers, err := datastore.GetSpeakersByIds(event.SpeakerIds)
		if err != nil {
			log.Printf("Failed to get speakers: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		event.Speakers = *speakers

		// Marshal the event.
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