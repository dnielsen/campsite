package handler

import (
	"campsite/services/session/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// `/{id}/comments` POST route.
func CreateComment(api service.CommentAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		sessionId := vars[ID]

		// Decode the body.
		var i service.CommentInput
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			log.Printf("Failed to unmarshal comment input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Create the comment in the database.
		c, err := api.CreateComment(sessionId, i)
		if err != nil {
			log.Printf("Failed to create comment: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Marshal the comment.
		commentBytes, err := json.Marshal(c)
		if err != nil {
			log.Printf("Failed to marshal comment: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond JSON with the session.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(commentBytes)
	}
}
