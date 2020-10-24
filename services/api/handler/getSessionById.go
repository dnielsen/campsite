package handler

import (
	"fmt"
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

// `/sessions/{id}` GET route. It communicates with the session service only.
func GetSessionById(c *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]
		// Create the request that calls our session service to get it.
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%v:%v/%v", c.Service.Session.Host, c.Service.Session.Port, id), nil)
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
		// Read the response body (hopefully it's our session).
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
