package handler

import (
	"dave-web-app/packages/server/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func DeleteSession(datastore service.SessionDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]

		// Delete the session from the database.
		if err := datastore.DeleteSession(id); err != nil {
			log.Printf("Failed to delete session: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond that the session deletion has been successful.
		w.WriteHeader(http.StatusNoContent)
	}
}