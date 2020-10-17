package handler

import (
	"campsite/pkg/model"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// `/sessions/{id}` DELETE route.
func DeleteSessionById(api model.SessionAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Currently our database doesn't know about `User` entity
		// so we're just ignoring claims.
		if _, err := verifyToken(r); err != nil {
			log.Printf("Failed to verify token: %v", err)
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]
		// Request the session service to delete the session from the database.
		if err := api.DeleteSessionById(id); err != nil {
			log.Printf("Failed to delete session: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Respond that the session has been successfully deleted.
		w.WriteHeader(http.StatusNoContent)
	}
}
