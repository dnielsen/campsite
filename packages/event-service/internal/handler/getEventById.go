package handler

import (
	"dave-web-app/packages/event-service/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const EVENT_ID = "eventId"

func GetEventById(datastore service.Datastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars[EVENT_ID]
		event, err := datastore.GetEventById(id)
		if err != nil {
			log.Printf("Failed to get event: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sessions, err := datastore.GetSessionsByEventId(id)
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

		for i, session := range event.Sessions {
			event.Sessions[i].Speakers = mapSpeakersToSpeakerIds(event.Speakers, session.SpeakerIds)
		}

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

func mapSpeakersToSpeakerIds(eventSpeakers []service.Speaker, sessionSpeakerIds []string) []service.Speaker {
	var sessionSpeakers []service.Speaker
	for _, sessionSpeakerId := range sessionSpeakerIds {
		for _, eventSpeaker := range eventSpeakers {
			if eventSpeaker.ID == sessionSpeakerId {
				sessionSpeakers = append(sessionSpeakers, eventSpeaker)
				break
			}
		}
	}
	return sessionSpeakers
}