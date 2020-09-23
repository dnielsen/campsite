package handler

import (
	"dave-web-app/packages/event-service/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func DeleteEventById(datastore service.EventAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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