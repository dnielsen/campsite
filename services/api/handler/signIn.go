package handler

import (
	"encoding/json"
	"github.com/dnielsen/campsite/pkg/model"
	"github.com/dnielsen/campsite/services/api/service"
	"log"
	"net/http"
)

// `/auth/sign-in` POST route. It communicates with the event service only.
func SignIn(api service.AuthAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the body.
		var i model.SignInInput
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			log.Printf("Failed to unmarshal sign in input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Request the auth service to sign in.
		token, err := api.SignIn(i)
		if err != nil {
			log.Printf("Failed to sign in: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("token:%v err:%v", token, err)
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