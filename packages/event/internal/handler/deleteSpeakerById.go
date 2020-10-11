package handler

import (
	"campsite/packages/event/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func DeleteSpeakerById(api service.SpeakerAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Currently our database doesn't know about `User` entity
		// so we're just ignoring claims.
		if _, err := verifyToken(r); err != nil {
			log.Printf("Failed to verify token: %v", err)
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]

		// Request the speaker service to delete the speaker from the database.
		if err := api.DeleteSpeakerById(id); err != nil {
			log.Printf("Failed to delete speaker: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Respond that the speaker has been successfully deleted.
		w.WriteHeader(http.StatusNoContent)
	}
}