package handler

import (
	"fmt"
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/dnielsen/campsite/pkg/jwt"
	"io/ioutil"
	"log"
	"net/http"
)

// `/speakers` POST route. It's a protected route. It communicates with the speakers service only.
func CreateSpeaker(c *config.Config) http.HandlerFunc {
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
		// Create the request that's gonna call our speaker service.
		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%v:%v", c.Service.Speaker.Host, c.Service.Speaker.Port), r.Body)
		if err != nil {
			log.Printf("Failed to create new request: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Make the request.
		//zipkinHttp.NewClient(t, zipkinHttp.ClientTrace(true))
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("Failed to do request: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Read the response body (the created speaker if the request was successful).
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Printf("Failed to read response body: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond with the received response (the status code is hopefully 201 Status Created).
		w.Header().Set(CONTENT_TYPE, res.Header.Get(CONTENT_TYPE))
		w.WriteHeader(res.StatusCode)
		w.Write(b)
	}
}
