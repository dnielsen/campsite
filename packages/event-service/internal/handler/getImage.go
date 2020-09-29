package handler

import (
	"campsite/packages/event-service/internal/service"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func GetImage(api service.ImageAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the filename parameter.
		vars := mux.Vars(r)
		filename := vars[FILENAME]

		// Get the image.
		img, err := api.GetImage(filename)
		if err != nil {
			log.Printf("Failed to get image: %v", err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		defer img.Close()

		// Respond with the image.
		w.Header().Set("Content-Type", "image/jpeg")
		w.WriteHeader(http.StatusOK)
		io.Copy(w, img)
	}
}