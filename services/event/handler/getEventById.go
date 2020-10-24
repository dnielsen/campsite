package handler

import (
	"encoding/json"
	"github.com/dnielsen/campsite/services/event/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// `/events/{id}` GET route. It communicates with the database only.
func GetEventById(api service.EventAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]
		// Get the event from the database.
		event, err := api.GetEventById(id)
		if err != nil {
			log.Printf("Failed to get event: %v", err)
			http.NotFound(w, r)
			return
		}
		// Add the (unique) speakers property to the event for our <FullEvent />
		// so that we don't need to do it on the frontend.
		event.Speakers = getUniqueSpeakersFromSessions(event.Sessions)
		// Marshal the event.
		b, err := json.Marshal(event)
		if err != nil {
			log.Printf("Failed to marshal event: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond JSON with the event
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
