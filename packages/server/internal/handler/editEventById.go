package handler

import (
	"campsite/packages/server/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// `/events/{id}` PUT route. It communicates with the database only.
func EditEventById(api service.EventAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Verify that the user is logged in and get the claims with the user email.
		claims, err := api.VerifyToken(r)
		if err != nil {
			log.Printf("Failed to verify token: %v", err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
		// Check permissions.
		_, err = api.VerifyRole(claims.ID, ADMIN_ONLY_ROLE_WHITELIST)
		if err != nil {
			log.Printf("Failed to verify permissions: %v", err)
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