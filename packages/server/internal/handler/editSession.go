package handler

import (
	"dave-web-app/packages/server/internal/service"
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func EditSession(datastore service.SessionDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]

		// Decode the body.
		var i service.SessionInput
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			log.Printf("Failed to unmarshal speaker input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Validate the input.
		log.Println(i)
		if err := validator.New().Struct(i); err != nil {
			log.Printf("Failed to validate session input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Update the session in the database.
		if err := datastore.EditSession(id, i); err != nil {
			log.Printf("Failed to edit session: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond that the session edit has been successful.
		w.WriteHeader(http.StatusOK)
	}
}
