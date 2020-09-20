package handler

import (
	"dave-web-app/packages/event-service/internal/service"
	"dave-web-app/packages/event-service/internal/util"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetEventById(datastore service.Datastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		sessions, err := datastore.GetSessionsByEventId(id)
		if err != nil {
			log.Printf("Failed to get sessions: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		event.Sessions = *sessions
		log.Println(event.Sessions)

		var speakerIds []string

		for _, session := range event.Sessions {
			speakerIds = append(speakerIds, session.SpeakerIds...)
		}
		uniqSpeakerIds := util.StrUnique(speakerIds)

		// Get the speakers from the speaker service.
		speakers, err := datastore.GetSpeakersByIds(uniqSpeakerIds)
		if err != nil {
			log.Printf("Failed to get speakers: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		event.Speakers = *speakers

		for i, session := range event.Sessions {
			for _, speaker := range *speakers {
				if util.StrContains(session.SpeakerIds, speaker.ID) {
					event.Sessions[i].Speakers = append(event.Sessions[i].Speakers, speaker)
				}
			}
		}

		// Marshal the event.
		eventBytes, err := json.Marshal(event)
		if err != nil {
			log.Printf("Failed to marshal full event: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(eventBytes)
	}
}