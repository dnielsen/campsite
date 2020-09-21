package handler

import (
	"dave-web-app/packages/event-service/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func DeleteSpeakerById(datastore service.SpeakerDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]

		// Request the speaker service to delete the speaker from the database.
		if err := datastore.DeleteSpeakerById(id); err != nil {
			log.Printf("Failed to delete speaker: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Respond that the speaker has been successfully deleted.
		w.WriteHeader(http.StatusNoContent)
	}
}