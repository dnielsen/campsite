package handler

import (
	"dave-web-app/packages/event-service/internal/service"
	"dave-web-app/packages/event-service/internal/util"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

const ID = "id"

func GetEventById(datastore service.Datastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		strId := vars[ID]
		intId, err := strconv.Atoi(strId)
		if err != nil {
			log.Printf("Failed to convert string id to integer")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id := uint(intId)
		event, err := datastore.GetEventById(id)
		if err != nil {
			log.Printf("Failed to get event: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sessions, err := datastore.GetSessionsByIds(event.SessionIds)
		if err != nil {
			log.Printf("Failed to get sessions: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		speakers, err := datastore.GetSpeakersByIds(event.SpeakerIds)
		if err != nil {
			log.Printf("Failed to get speakers: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
