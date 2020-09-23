package handler

import (
	"dave-web-app/packages/server/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// `/events/{id}` GET route.
func GetEventById(api service.EventAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]

		// Get the event from the database.
		event, err := api.GetEventById(id)
		if err != nil {
			log.Printf("Failed to get event: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Marshal the event.
		eventBytes, err := json.Marshal(event)
		if err != nil {
			log.Printf("Failed to marshal full event: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond JSON with the events.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(eventBytes)
	}
}
