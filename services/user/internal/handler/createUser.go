package handler

import (
	"campsite/packages/user/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

// `/` POST route. On successful creation it returns the created user and the code 201 (Status Created).
func CreateUser(api service.UserAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the body.
		var i service.CreateUserInput
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			log.Printf("Failed to unmarshal create user input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Create the user in the database
		u, err := api.CreateUser(i)
		if err != nil {
			log.Printf("Failed to create user: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Marshal the user.
		b, err := json.Marshal(u)
		if err != nil {
			log.Printf("Failed to marshal user: %v", err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
		// Respond JSON with the user.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(b)
	}
}