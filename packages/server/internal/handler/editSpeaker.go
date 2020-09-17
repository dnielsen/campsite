
package handler

import (
	"dave-web-app/packages/server/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func EditSpeaker(datastore service.SpeakerDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]

		// Decode the body.
		var i service.SpeakerInput
		err := json.NewDecoder(r.Body).Decode(&i)
		if err != nil {
			log.Printf("Failed to unmarshal speaker input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Update the speaker in the database.
		speaker, err := datastore.EditSpeaker(id, i)
		if err != nil {
			log.Printf("Failed to edit speaker: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Marshal the updated speaker.
		speakerBytes, err := json.Marshal(speaker)
		if err != nil {
			log.Printf("Failed to marshal speaker: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond JSON with the updated speaker.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(speakerBytes)
	}
}