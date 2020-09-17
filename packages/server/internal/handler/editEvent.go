package handler

import (
	"dave-web-app/packages/server/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func EditEvent(datastore service.EventDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]

		// Decode the body.
		var i service.EventInput
		err := json.NewDecoder(r.Body).Decode(&i)
		if err != nil {
			log.Printf("Failed to unmarshal event input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Update the event in the database.
		event, err := datastore.EditEvent(id, i)
		if err != nil {
			log.Printf("Failed to edit event: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Marshal the bytes.
		eventBytes, err := json.Marshal(event)
		if err != nil {
			log.Printf("Failed to marshal event: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond JSON with the updated event.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(eventBytes)
	}
}