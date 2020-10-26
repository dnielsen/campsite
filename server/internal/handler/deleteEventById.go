package handler

import (
	"campsite/packages/server/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// `/events/{id}` DELETE route.
func DeleteEventById(api service.EventAPI) http.HandlerFunc {
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
		eventId := vars[ID]
		// Delete the event from the database.
		if err := api.DeleteEventById(eventId); err != nil {
			log.Printf("Failed to delete event: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Respond that the event has been successfully deleted.
		w.WriteHeader(http.StatusNoContent)
	}
}