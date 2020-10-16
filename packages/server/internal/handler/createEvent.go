package handler

import (
	"campsite/packages/server/internal/service"
	"campsite/packages/server/internal/service/role"
	"encoding/json"
	"log"
	"net/http"
)

// `/events` POST route.
func CreateEvent(api service.EventAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := api.VerifyToken(r)
		if err != nil {
			log.Printf("Failed to verify token: %v", err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		u, err := api.GetUserByEmail(claims.Email)
		if err != nil {
			log.Printf("Failed to get user by email: %v", err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if u.Role != role.ADMIN {
			log.Printf("Failed to verify role: needed role: %v, current role: %v", role.ADMIN, u.Role)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		// Decode the body.
		var i service.EventInput
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			log.Printf("Failed to unmarshal event input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Create the event in the database.
		event, err := api.CreateEvent(u.ID, i)
		if err != nil {
			log.Printf("Failed to create event: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Marshal the event.
		eventBytes, err := json.Marshal(event)
		if err != nil {
			log.Printf("Failed to marshal event: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond JSON with the created event
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(eventBytes)
	}
}
