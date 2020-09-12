package handler

import (
	"dave-web-app/packages/event-service/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const SESSION_ID = "sessionId"

func GetSessionById(datastore service.Datastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars[SESSION_ID]
		session, err := datastore.GetSessionById(id)
		if err != nil {
			log.Printf("Failed to get session: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		speakers, err := datastore.GetSpeakersByIds(session.SpeakerIds)
		log.Println(session.SpeakerIds)
		if err != nil {
			log.Printf("Failed to get speakers: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		session.Speakers = *speakers

		sessionBytes, err := json.Marshal(session)
		if err != nil {
			log.Printf("Failed to marshal session: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(sessionBytes)
	}
}
