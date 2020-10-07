package handler

import (
	"campsite/packages/event-service/internal/service"
	"encoding/json"
	"log"
	"net/http"
)




// We'll later move it to an environment variable.
var JWT_SECRET_KEY = []byte("V3RY_S3CR3T_K3Y")

func SignIn(api service.EventAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the body.
		var i service.SignInInput
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			log.Printf("Failed to unmarshal sign in input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Validate that the credentials match some user.
		u, err := api.ValidateUser(i)
		if err != nil {
			log.Printf("Failed to sign in: %v", err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := generateToken(u.Email)
		if err != nil {
			log.Printf("Failed to generate token: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//setTokenCookie(w, token)

		w.Write([]byte(token))
		// Respond that the sign in has been successful.
	}
}