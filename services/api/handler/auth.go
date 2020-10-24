package handler

import (
	"encoding/json"
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/dnielsen/campsite/pkg/jwt"
	"log"
	"net/http"
)

// `/auth` GET route. It communicates with the auth service only.
func Auth(c *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Try to get claims from the token if exists.
		tokenCookie, err := r.Cookie(c.Jwt.CookieName)
		if err != nil {
			log.Printf("Failed to get cookie: %v", err)
			// err != nil == no cookie == not signed in
			w.WriteHeader(http.StatusOK)
			return
		}
		claims, err := jwt.VerifyToken(tokenCookie.Value, &c.Jwt)
		if err != nil {
			log.Printf("Failed to verify token: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Marshal the `Me`.
		b, err := json.Marshal(claims.Me)
		if err != nil {
			log.Printf("Failed to marshal claims: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond JSON with the auth data (`Me`)
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}