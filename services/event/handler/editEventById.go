package handler

import (
	"encoding/json"
	"github.com/dnielsen/campsite/pkg/model"
	"github.com/dnielsen/campsite/services/event/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// `/events/{id}` PUT route. It communicates with the database only.
func EditEventById(api service.EventAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]
		// Decode the body.
		var i model.EventInput
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			log.Printf("Failed to unmarshal event input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Edit the event in the database.
		e, err := api.EditEventById(id, i)
		if err != nil {
			log.Printf("Failed to edit event: %v", err)
			http.NotFound(w, r)
			return
		}
		b, err := json.Marshal(e)
		if err != nil {
			log.Printf("Failed to marshal event: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond with the edited event.
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func getUniqueSpeakersFromSessions(sessions []model.Session) []model.Speaker {
	// Iterate through each session and each speaker of the sessions
	// and create an array of unique speakers.
	// The key of the map is a speaker id and the value is a speaker.
	spkMap := make(map[string]model.Speaker)
	for _, sess := range sessions {
		for _, spk := range sess.Speakers {
			spkMap[spk.ID] = spk
		}
	}

	// Get the values (Speakers) of the map.
	var uniqSpks []model.Speaker
	for _, spk := range spkMap {
		uniqSpks = append(uniqSpks, spk)
	}

	return uniqSpks
}
