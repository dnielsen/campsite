package handler

import (
	"fmt"
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/dnielsen/campsite/pkg/jwt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// `/speakers/{id}` DELETE route. It's a protected route. It communicates with the speaker service only.
func DeleteSpeakerById(client *http.Client, c *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Verify the JWT token since it's a protected route.
		tokenCookie, err := r.Cookie(c.Jwt.CookieName)
		if err != nil {
			log.Printf("Failed to get cookie: %v", err)
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		_, err = jwt.VerifyToken(tokenCookie.Value, &c.Jwt)
		if err != nil {
			log.Printf("Failed to verify token: %v", err)
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]
		// Create the request that calls our speaker service to delete it.
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%v:%v/%v", c.Service.Speaker.Host, c.Service.Speaker.Port, id), nil)
		if err != nil {
			log.Printf("Failed to create new request: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Make the request.
		res, err := client.Do(req)
		if err != nil {
			log.Printf("Failed to do request: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond with the received response (hopefully it's 204 No Content).
		w.WriteHeader(res.StatusCode)
	}
}
