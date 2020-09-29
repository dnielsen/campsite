package handler

import (
	"campsite/packages/event-service/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// `/sessions/{id}` PUT route.
func EditSessionById(api service.SessionAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]
		// Decode the body.
		var i service.SessionInput
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			log.Printf("Failed to unmarshal session input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Edit the edit in the database.
		if err := api.EditSessionById(id, i); err != nil {
			log.Printf("Failed to edit session: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond that the session has been edited successfully.
		w.WriteHeader(http.StatusNoContent)
	}
}