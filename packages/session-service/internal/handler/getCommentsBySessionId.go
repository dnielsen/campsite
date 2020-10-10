package handler

import (
	"campsite/packages/session-service/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

const (
	LIMIT = "limit"
	CURSOR = "cursor"
)

type CommentsResponse struct {
	Comments *[]service.Comment `json:"comments"`
	EndCursor *string `json:"endCursor"`
}

func GetCommentsBySessionId(datastore service.CommentAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		// Get the id parameter.
		vars := mux.Vars(r)
		sessionId := vars[ID]

		// If there's no limit parameter specified by default it's gonna be 5.
		limit := 5

		strLimit := r.URL.Query().Get(LIMIT)
		if strLimit != "" {
			limit, err = strconv.Atoi(strLimit)
			if err != nil || limit < 0 {
				http.Error(w, "limit parameter invalid: it must be a positive integer", http.StatusBadRequest)
				return
			}
		}

		// The cursor will get validated later.
		cursor := r.URL.Query().Get(CURSOR)

		// Get the comments from the database.
		comments, endCursor, err := datastore.GetCommentsBySessionId(sessionId, limit, cursor)
		if err != nil {
			log.Printf("Failed to get comments: %v", err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		resData := CommentsResponse{
			Comments:  comments,
			EndCursor: endCursor,
		}

		// Marshal the comments.
		commentBytes, err := json.Marshal(resData)
		if err != nil {
			log.Printf("Failed to marshal comments data: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond JSON with the comments.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(commentBytes)
	}
}
