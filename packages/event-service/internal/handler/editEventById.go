package handler

import (
	"campsite/packages/event-service/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// `/events/{id}` PUT route.
func EditEventById(api service.EventAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Currently our database doesn't know about `User` entity
		// so we're just ignoring claims.
		_, err := verifyToken(w, r)
		if err != nil {
			log.Printf("Failed to verify token: %v", err)
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]
		// Decode the body.
		var i service.EventInput
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			log.Printf("Failed to unmarshal event input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Edit the event in the database.
		if err := api.EditEventById(id, i); err != nil {
			log.Printf("Failed to edit event: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond that the event has been edited successfully.
		w.WriteHeader(http.StatusNoContent)
	}
}