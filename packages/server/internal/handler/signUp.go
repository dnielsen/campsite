package handler

import (
	"campsite/packages/server/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

// `/sign-up` POST route. On successful sign up it returns a JWT token and the code 201 (Status Created).
func SignUp(api service.UserAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the body.
		var i service.SignUpInput
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			log.Printf("Failed to unmarshal sign up input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		u, err := api.CreateUser(i)
		if err != nil {
			log.Printf("Failed to sign up: %v", err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		log.Printf("User with the email %v and role %v has been created in the database!", u.Email, u.Role)
		// Generate the access token.
		token, err := api.GenerateToken(u)
		if err != nil {
			log.Printf("Failed to generate token: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond plain text with the token. We might change the response later,
		// to some json object
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(token))
	}
}