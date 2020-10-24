package handler

import (
	"encoding/json"
	"fmt"
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/dnielsen/campsite/pkg/jwt"
	"io/ioutil"
	"log"
	"net/http"
)

// `/auth/sign-in` POST route. It communicates with the auth service only.
func SignIn(c *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Create the request.
		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%v:%v/sign-in", c.Service.Auth.Host, c.Service.Auth.Port), r.Body)
		if err != nil {
			log.Printf("Failed to create new request: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Make the request.
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("Failed to do request: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Read the response body.
		tokenBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Printf("Failed to read response body: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Set the token cookie.
		token := string(tokenBytes)
		http.SetCookie(w, &http.Cookie{
			Name:       c.Jwt.CookieName,
			Value:      token,
			Path:       "/",
			MaxAge:     60 * 60 * 24 * 7,
			Secure:     false,
			HttpOnly:   true,
			SameSite:   0,
		})
		// Get the claims.
		claims, err := jwt.VerifyToken(token, &c.Jwt)
		if err != nil {
			log.Printf("Failed to verify token: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Marshal `Me` (auth data).
		b, err := json.Marshal(claims.Me)
		if err != nil {
			log.Printf("Failed to marshal `Me`: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond JSON with `Me`
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}