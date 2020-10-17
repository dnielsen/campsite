package handler

import (
	"campsite/pkg/model"
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
func GetCommentsBySessionId(api model.SessionAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		sessionId := vars[ID]
		// Get the query parameters.
		limit := r.URL.Query().Get(LIMIT)
		cursor := r.URL.Query().Get(CURSOR)
		// Get comments from the database.
		commentsData, err := api.GetCommentsBySessionId(sessionId, limit, cursor)
		if err != nil {
			log.Printf("Failed to get comments: %v", err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		// Marshal the comments.
		b, err := json.Marshal(commentsData)
		if err != nil {
			log.Printf("Failed to marshal comments data: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond JSON with the comments.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
