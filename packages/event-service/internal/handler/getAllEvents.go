package handler

import (
	"dave-web-app/packages/event-service/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

func GetAllEvents(datastore service.EventDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get all events from the database.
		events, err := datastore.GetAllEvents()
		if err != nil {
			log.Printf("Failed to get events: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Marshal the events.
		eventBytes, err := json.Marshal(events)
		if err != nil {
			log.Printf("Failed to marshal events: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond JSON with the events.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(eventBytes)
	}
}
