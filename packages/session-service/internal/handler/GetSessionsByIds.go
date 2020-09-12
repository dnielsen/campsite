package handler

import (
	"dave-web-app/packages/session-service/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

const ID = "id"

type getSessionsByIdsRequestBody struct {
	SessionIds []string `json:"sessionIds"`
}

func GetSessionsByIds(datastore service.SessionDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("hello session")
		var body getSessionsByIdsRequestBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			log.Printf("Failed to decode request body: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		log.Println(body.SessionIds)

		sessions, err := datastore.GetSessionsByIds(body.SessionIds)
		if err != nil {
			log.Printf("Failed to get sessions: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println(sessions)
		sessionBytes, err := json.Marshal(sessions)
		if err != nil {
			log.Printf("Failed to marshal sessions: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(sessionBytes)
	}
}
