package handler

import (
	"encoding/json"
	"github.com/dnielsen/campsite/pkg/model"
	"github.com/dnielsen/campsite/services/event/handler"
	"github.com/dnielsen/campsite/services/event/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// `/speakers/{id}` PUT route. It communicates with the speaker service only.
func EditSpeakerById(api service.SpeakerAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[handler.ID]
		// Decode the body.
		var i model.SpeakerInput
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			log.Printf("Failed to unmarshal speaker input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Edit the speaker in the database.
		if err := api.EditSpeakerById(id, i); err != nil {
			log.Printf("Failed to edit speaker: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond that the speaker has been edited successfully.
		w.WriteHeader(http.StatusNoContent)
	}
}