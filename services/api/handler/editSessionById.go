package handler

import (
	"encoding/json"
	"github.com/dnielsen/campsite/pkg/model"
	"github.com/dnielsen/campsite/services/api/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// `/sessions/{id}` PUT route. It communicates with the session service only.
func EditSessionById(api service.SessionAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]
		// Decode the body.
		var i model.SessionInput
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			log.Printf("Failed to unmarshal session input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Edit the session in the database.
		s, err := api.EditSessionById(id, i)
		if err != nil {
			log.Printf("Failed to edit session: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Marshal the session
		b, err := json.Marshal(s)
		if err != nil {
			log.Printf("Failed to marshal session: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond JSON with the session.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
