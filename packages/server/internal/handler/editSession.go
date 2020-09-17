
package handler

import (
	"dave-web-app/packages/server/internal/service"
	"encoding/json"
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
		err := json.NewDecoder(r.Body).Decode(&i)
		if err != nil {
			log.Printf("Failed to unmarshal speaker input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Update the session in the database.
		session, err := datastore.EditSession(id, i)
		if err != nil {
			log.Printf("Failed to edit session: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Marshal the bytes.
		sessionBytes, err := json.Marshal(session)
		if err != nil {
			log.Printf("Failed to marshal session: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond JSON with the updated speaker.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(sessionBytes)
	}
}