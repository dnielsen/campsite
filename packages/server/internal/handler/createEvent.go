
package handler

import (
	"dave-web-app/packages/server/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

func CreateEvent(datastore service.EventDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i service.EventInput

		// Decode the body
		err := json.NewDecoder(r.Body).Decode(&i)
		if err != nil {
			log.Printf("Failed to unmarshal event input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Create the event in the database.
		event, err := datastore.CreateEvent(i)
		if err != nil {
			log.Printf("Failed to create event: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		eventBytes, err := json.Marshal(event)
		if err != nil {
			log.Printf("Failed to marshal event: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(eventBytes)
	}
}