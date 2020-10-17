package handler

import (
	"campsite/packages/event/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func DeleteEventById(api service.EventAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Currently our database doesn't know about `User` entity
		// so we're just ignoring claims.
		if _, err := verifyToken(r); err != nil {
			log.Printf("Failed to verify token: %v", err)
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]

		// Delete the event from the database.
		if err := api.DeleteEventById(id); err != nil {
			log.Printf("Failed to delete event: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Respond that the event has been successfully deleted.
		w.WriteHeader(http.StatusNoContent)
	}
}
