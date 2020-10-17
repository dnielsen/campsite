package handler

import (
	"campsite/packages/server/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

// `/speakers` POST route.
func CreateSpeaker(api service.SpeakerAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Verify that the user is logged in and get the claims with the user email.
		claims, err := api.VerifyToken(r)
		if err != nil {
			log.Printf("Failed to verify token: %v", err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
		// Check permissions
		u, err := api.VerifyRole(claims.ID, ADMIN_ONLY_ROLE_WHITELIST)
		if err != nil {
			log.Printf("Failed to verify permissions: %v", err)
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		// Decode the body.
		var i service.SpeakerInput
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			log.Printf("Failed to unmarshal speaker input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Create the speaker in the database.
		speaker, err := api.CreateSpeaker(i, u.ID)
		if err != nil {
			log.Printf("Failed to create speaker: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Marshal the speaker.
		speakerBytes, err := json.Marshal(speaker)
		if err != nil {
			log.Printf("Failed to marshal speaker: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond JSON with the created speaker.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(speakerBytes)
	}
}