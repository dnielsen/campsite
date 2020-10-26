package handler

import (
	"encoding/json"
	"github.com/dnielsen/campsite/services/event/service"
	"log"
	"net/http"
)

// `/events` GET route. It communicates with the database only.
func GetAllEvents(api service.EventAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get all events from the database.
		events, err := api.GetAllEvents()
		if err != nil {
			log.Printf("Failed to get all events: %v", err)
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
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		w.WriteHeader(http.StatusOK)
		w.Write(eventBytes)
	}
}
