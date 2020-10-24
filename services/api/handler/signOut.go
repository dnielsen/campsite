package handler

import (
	"github.com/dnielsen/campsite/pkg/config"
	"net/http"
)

// `/auth/sign-out` POST route. It communicates with the auth service only.
func SignOut(c *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Reset the token cookie.
		http.SetCookie(w, &http.Cookie{
			Name:       c.Jwt.CookieName,
			Value:      "",
			Path:       "/",
			MaxAge:     -1,
			Secure:     false,
			HttpOnly:   true,
			SameSite:   0,
		})
		// Respond with the received status code.
		w.WriteHeader(http.StatusOK)
	}
}