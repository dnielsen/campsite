package handler

import (
	"encoding/json"
	"github.com/dnielsen/campsite/services/api/service"
	"log"
	"net/http"
)

// `/events` GET route. It communicates with the events service only.
func GetAllEvents(api service.EventAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get all events from the events service.
		events, err := api.GetAllEvents()
		if err != nil {
			log.Printf("Failed to get events: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Marshal the events.
		b, err := json.Marshal(events)
		if err != nil {
			log.Printf("Failed to marshal events: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond json with the events.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
