package handler

import (
	"campsite/services/event/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	LIMIT  = "limit"
	CURSOR = "cursor"
)

// `/sessions/{id}/comments` GET route.
func GetCommentsBySessionId(api service.SessionAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]

		limit := r.URL.Query().Get(LIMIT)
		cursor := r.URL.Query().Get(CURSOR)

		// Get comments from the database.
		commentsData, err := api.GetCommentsBySessionId(id, limit, cursor)
		if err != nil {
			log.Printf("Failed to get comments: %v", err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		// Marshal the comments.
		bytes, err := json.Marshal(commentsData)
		if err != nil {
			log.Printf("Failed to marshal comments data: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond JSON with the comments.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}
