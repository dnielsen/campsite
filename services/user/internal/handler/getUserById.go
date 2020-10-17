package handler

import (
	"campsite/services/user/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const ID = "id"

func GetUserById(api service.UserAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]
		// Grab the user from the database.
		u, err := api.GetUserById(id)
		if err != nil {
			log.Printf("Failed to get user by id: %v", err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		// Marshal the user.
		b, err := json.Marshal(u)
		if err != nil {
			log.Printf("Failed to marshal user: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond JSON with the user.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}