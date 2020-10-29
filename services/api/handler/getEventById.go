package handler

import (
	"fmt"
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

// `/events/{id}` GET route. It communicates with the event service only.
func GetEventById(client *http.Client, c *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]
		// Create the request that calls our event service to get it.
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%v:%v/%v", c.Service.Event.Host, c.Service.Event.Port, id), nil)
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
		// Read the response body (hopefully our event).
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Printf("Failed to read body: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond with the received response.
		w.Header().Set(CONTENT_TYPE, r.Header.Get(CONTENT_TYPE))
		w.WriteHeader(res.StatusCode)
		w.Write(b)
	}
}
