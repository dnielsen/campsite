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
		// Get the request parameters (id).
		vars := mux.Vars(r)
		id := vars[SPEAKER_ID]

		// Get the speaker from the database
		speaker, err := datastore.GetSpeakerById(id)
		if err != nil {
			log.Printf("Failed to get speaker: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println(speaker.SessionIds)
		// Get the sessions from the database to attach it to the speaker struct
		sessions, err := datastore.GetSessionsByIds(speaker.SessionIds)
		if err != nil {
			log.Printf("Failed to get sessions: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println(sessions)
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
