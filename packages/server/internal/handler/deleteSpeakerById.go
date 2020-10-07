package handler

import (
	"campsite/packages/server/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// `/speakers/{id}` DELETE route.
func DeleteSpeakerById(api service.SpeakerAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]

		// Delete the speaker from the database.
		if err := api.DeleteSpeakerById(id); err != nil {
			log.Printf("Failed to delete speaker: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Respond that the speaker has been successfully deleted.
		w.WriteHeader(http.StatusNoContent)
	}
}