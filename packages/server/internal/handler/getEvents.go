package handler

import (
	"dave-web-app/packages/server/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

// `/events` GET route.
func GetEvents(api service.EventAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// We can later add url parameters to
		// specify the ids of the events we want
		// if needed.

		// Get the events from the database.
		events, err := api.GetAllEvents()
		if err != nil {
			log.Printf("Failed to get events: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Marshal the event.
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
