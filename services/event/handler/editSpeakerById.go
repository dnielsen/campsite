package handler

import (
	"campsite/pkg/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// `/speaker/{id}` PUT route.
func EditSpeakerById(api model.SpeakerAPI) http.HandlerFunc {
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
		// Decode the body.
		var i model.SpeakerInput
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			log.Printf("Failed to unmarshal speaker input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Edit the speaker in the database.
		s, err := api.EditSpeakerById(id, i)
		if err != nil {
			log.Printf("Failed to edit speaker: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Marshal the speaker
		b, err := json.Marshal(s)
		if err != nil {
			log.Printf("Failed to marshal speaker: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond that the speaker has been edited successfully.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		w.Write(b)
	}
}
