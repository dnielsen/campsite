package handler

import (
	"dave-web-app/packages/server/internal/service"
	"encoding/json"
	"github.com/go-playground/validator"
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

		// Validate the input.
		if err := validator.New().Struct(i); err != nil {
			log.Printf("Failed to validate speaker input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Update the speaker in the database.
		if err := datastore.EditSpeaker(id, i); err != nil {
			log.Printf("Failed to edit speaker: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond the speaker edit has been successful.
		w.WriteHeader(http.StatusOK)
	}
}
