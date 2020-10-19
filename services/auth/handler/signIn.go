package handler

import (
	"encoding/json"
	"github.com/dnielsen/campsite/pkg/model"
	"log"
	"net/http"
	"time"
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
		// Validate the credentials match the user.
		u, err := api.ValidateUser(i)
		if err != nil {
			log.Printf("Failed to validate user: %v", err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// Generate the JWT token.
		token, err := api.GenerateToken(u)
		if err != nil {
			log.Printf("Failed to generate token")
		}
		// Send the token with the cookies.
		http.SetCookie(w, &http.Cookie{
			Name:       "token",
			Value:      token,
			Path:       "/",
			Expires:    time.Time{},
			RawExpires: "",
			MaxAge:     60 * 60 * 24 * 7,
			Secure:     false,
			HttpOnly:   true,
			SameSite:   0,
		})
		w.WriteHeader(http.StatusOK)
	}
}
