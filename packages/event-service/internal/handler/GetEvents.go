package handler

import (
	"dave-web-app/packages/event-service/internal/service"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// the `/` route. You could specify an id parameter like so: `/?id=453,123,435` which would fetch the
// speakers with the ids of 453, 123, and 435.
func GetEvents(datastore service.Datastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the ids from the query parameter.
		idsString := r.URL.Query().Get(ID)
		ids := strings.Split(idsString, ",")

		// Get the events from the database.
		var err error
		var events *[]service.Event
		// If there are no parameters, by default there's a string array with an empty string inside.
		// So, we're checking for an array of length 1 with an empty string inside.
		idsNotSpecified := len(ids) == 1 && len(ids[0]) == 0
		if idsNotSpecified {
			events, err = datastore.GetAllEvents()
		} else {
			//sessions, err = datastore.GetEventsByIds(ids)
			panic("get events by ids not implemented yet")
		}

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
