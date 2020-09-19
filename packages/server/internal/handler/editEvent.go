package handler

import (
	"dave-web-app/packages/server/internal/service"
	"encoding/json"
	"github.com/go-playground/validator"
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

		// Validate the input.
		if err := validator.New().Struct(i); err != nil {
			log.Printf("Failed to validate event input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Update the event in the database.
		if err := datastore.EditEvent(id, i); err != nil {
			log.Printf("Failed to edit event: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond that the event edit has been successful.
		w.WriteHeader(http.StatusOK)
	}
}
