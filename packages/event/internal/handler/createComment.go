package handler

import (
	"campsite/packages/event/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// `/sessions/{id}/comments` POST route.
func CreateComment(api service.SessionAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Currently our database doesn't know about `User` entity
		// so we're just ignoring claims.
		if _, err := verifyToken(r); err != nil {
			log.Printf("Failed to verify token: %v", err)
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		// Get the session id parameter.
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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Marshal the comment.
		commentBytes, err := json.Marshal(c)
		if err != nil {
			log.Printf("Failed to marshal comm: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond JSON with the created event
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(commentBytes)
	}
}