package handler

import (
	"campsite/packages/event-service/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func DeleteEventById(datastore service.EventAPI) http.HandlerFunc {
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

		// Delete the event from the database.
		if err := datastore.DeleteEventById(id); err != nil {
			log.Printf("Failed to delete event: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Respond that the event has been successfully deleted.
		w.WriteHeader(http.StatusNoContent)
	}
}