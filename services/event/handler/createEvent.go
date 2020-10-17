package handler

import (
	"campsite/services/event/service"
	"encoding/json"
	"log"
	"net/http"
)

func CreateEvent(api service.EventAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Currently our database doesn't know about `User` entity
		// so we're just ignoring claims.
		if _, err := verifyToken(r); err != nil {
			log.Printf("Failed to verify token: %v", err)
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		// Decode the body.
		var i service.EventInput
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			log.Printf("Failed to unmarshal event input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Create the event in the database.
		event, err := api.CreateEvent(i)
		if err != nil {
			log.Printf("Failed to create event: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Marshal the event.
		eventBytes, err := json.Marshal(event)
		if err != nil {
			log.Printf("Failed to marshal event: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond JSON with the created event
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(eventBytes)
	}
}
