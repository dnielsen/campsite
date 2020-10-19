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
		// Validate the credentials and generate the JWT token.
		token, err := api.SignIn(i)
		if err != nil {
			log.Printf("Failed to sign in: %v", err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Send the cookie with the token.
		http.SetCookie(w, &http.Cookie{
			Name:       "token",
			Value:      token,
			Path:       "/",
			MaxAge:     60 * 60 * 24 * 7,
			Secure:     false,
			HttpOnly:   true,
			SameSite:   0,
		})
		w.WriteHeader(http.StatusOK)
	}
}
