package handler

import (
	"fmt"
	"github.com/dnielsen/campsite/pkg/config"
	"io/ioutil"
	"log"
	"net/http"
)

// `/events` GET route. It communicates with the event service only.
func GetAllEvents(c *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Create the request that calls our event service to get the events.
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%v:%v", c.Service.Event.Host, c.Service.Event.Port), nil)
		if err != nil {
			log.Printf("Failed to create new request: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Make the request.
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("Failed to do request: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Read the response body (hopefully it's our events).
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
