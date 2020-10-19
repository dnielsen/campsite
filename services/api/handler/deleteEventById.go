package handler

import (
	"github.com/dnielsen/campsite/services/api/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// `/events/{id}` DELETE route. It communicates with the event service only.
func DeleteEventById(api service.EventAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]
		// Request the event service to delete the event from the database.
		if err := api.DeleteEventById(id); err != nil {
			log.Printf("Failed to delete session: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Respond that the event has been successfully deleted.
		w.WriteHeader(http.StatusNoContent)
	}
}
