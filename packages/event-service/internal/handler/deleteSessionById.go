package handler

import (
	"dave-web-app/packages/event-service/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func DeleteSessionById(datastore service.SessionDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]

		// Request the session service to delete the session from the database.
		if err := datastore.DeleteSessionById(id); err != nil {
			log.Printf("Failed to delete session: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Respond that the session has been successfully deleted.
		w.WriteHeader(http.StatusNoContent)
	}
}