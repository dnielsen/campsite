package handler

import (
	"campsite/services/event/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// `/speakers/{id}` DELETE route. It communicates with the speaker service only.
func DeleteSpeakerById(datastore service.SpeakerAPI) http.HandlerFunc {
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
