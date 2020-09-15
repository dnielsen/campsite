package handler

import (
	"dave-web-app/packages/speaker-service/internal/service"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// the `/` route. You could specify id an parameter like so: `/?id=453,123,435` which would fetch the
// speakers with the ids of 453, 123, and 435.
func GetSpeakers(datastore service.Datastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the ids from the query parameter.
		idsString := r.URL.Query().Get(ID)
		ids := strings.Split(idsString, ",")

		// Get the speakers from the database.
		var err error
		var speakers *[]service.Speaker
		// If there are no parameters, by default there's a string array with an empty string inside.
		// So, we're checking for an array of length 1 with an empty string inside.
		idsNotSpecified := len(ids) == 1 && len(ids[0]) == 0
		if idsNotSpecified {
			speakers, err = datastore.GetAllSpeakers()
		} else {
			speakers, err = datastore.GetSpeakersByIds(ids)
		}

		if err != nil {
			log.Printf("Failed to get speakers: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Marshal the speakers.
		speakerBytes, err := json.Marshal(speakers)
		if err != nil {
			log.Printf("Failed to marshal speakers: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond json with the speakers.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(speakerBytes)
	}
}
