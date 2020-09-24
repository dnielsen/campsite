package handler

import (
	"dave-web-app/packages/server/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// `/sessions/{id}` DELETE route.
func DeleteSessionById(api service.SessionAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]

		// Delete the session from the database.
		if err := api.DeleteSessionById(id); err != nil {
			log.Printf("Failed to delete session: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond the session deletion has been successful.
		w.WriteHeader(http.StatusNoContent)
	}
}
