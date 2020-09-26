package handler

import (
	"dave-web-app/packages/server/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// `/events/{id}` DELETE route.
func DeleteEventById(api service.EventAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]

		// Delete the event from the database.
		if err := api.DeleteEventById(id); err != nil {
			log.Printf("Failed to delete event: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond the event deletion has been successful.
		w.WriteHeader(http.StatusNoContent)
	}
}