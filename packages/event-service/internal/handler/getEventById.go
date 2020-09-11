package handler

import (
	"dave-web-app/packages/event-service/internal/service"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const ID = "id"

func GetEventById(datastore service.EventDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		strId := vars[ID]
		id, err := uuid.Parse(strId)
		if err != nil {
			log.Printf("Failed to convert id to uuid: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		event, err := datastore.GetEventById(id)
		if err != nil {
			log.Printf("Failed to get event: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		eventBytes, err := json.Marshal(event)
		if err != nil {
			log.Printf("Failed to marshal event: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(eventBytes)
	}
}
