package handler

import (
	"encoding/json"
	"github.com/dnielsen/campsite/pkg/model"
	"github.com/dnielsen/campsite/services/auth/service"
	"log"
	"net/http"
)

// `/sign-in` POST route. It communicates with the database only.
func SignIn(api service.AuthAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the body.
		var i model.SignInInput
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			log.Printf("Failed to unmarshal sign in input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Validate the credentials, generate the JWT token, and add it to the header.
		token, err := api.SignIn(i)
		if err != nil {
			log.Printf("Failed to sign in: %v", err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// Respond plain text with the token.
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(token))
	}
}
