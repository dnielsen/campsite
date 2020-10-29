package handler

import (
	"fmt"
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/dnielsen/campsite/pkg/jwt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

// `/events/{id}` PUT route. It's a protected route. It communicates with the event service only.
func EditEventById(client *http.Client, c *config.Config) http.HandlerFunc {
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
		// Create the request that calls our event service to edit it.
		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://%v:%v/%v", c.Service.Event.Host, c.Service.Event.Port, id), r.Body)
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
		// Read the response body (hopefully our edited event).
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Printf("Failed to read body: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond with the received response (hopefully it's 200 Status OK).
		w.Header().Set(CONTENT_TYPE, r.Header.Get(CONTENT_TYPE))
		w.WriteHeader(res.StatusCode)
		w.Write(b)
	}
}
