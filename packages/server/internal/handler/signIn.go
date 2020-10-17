package handler

import (
	"campsite/packages/server/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

func SignIn(api service.AuthAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the body.
		var i service.SignInInput
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			log.Printf("Failed to unmarshal sign in input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Validate the credentials match the user.
		u, err := api.ValidateUser(i)
		if err != nil {
			log.Printf("Failed to authenticate: %v", err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// Generate the access token.
		token, err := api.GenerateToken(u)
		if err != nil {
			log.Printf("Failed to generate token: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond plain text with the token. We might change the response later,
		// to some json object
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(token))
	}
}