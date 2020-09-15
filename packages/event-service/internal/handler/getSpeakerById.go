package handler

import (
	"dave-web-app/packages/event-service/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetSpeakerById(datastore service.Datastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id paramter.
		vars := mux.Vars(r)
		id := vars[ID]

		// Get the speaker from the speaker service.
		speaker, err := datastore.GetSpeakerById(id)
		if err != nil {
			log.Printf("Failed to get speaker: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Get the sessions from the session service to attach it to the speaker
		// since we're always displaying the session for the full speaker item.
		sessions, err := datastore.GetSessionsByIds(speaker.SessionIds)
		if err != nil {
			log.Printf("Failed to get sessions: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		speaker.Sessions = *sessions

		// Marshal the speaker.
		speakerBytes, err := json.Marshal(speaker)
		if err != nil {
			log.Printf("Failed to marshal speaker: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond json with the found speaker.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(speakerBytes)
	}
}
