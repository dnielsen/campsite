package handler

import (
	"encoding/json"
	"github.com/dnielsen/campsite/services/event/service"
	"log"
	"net/http"
)

// Our frontend appends a file and sets the form data name to `file`.
// It's the most commonly used name for form data.
const FORM_DATA_NAME = "file"

// `/images` POST route. It doesn't communicate with the database or any of the services.
// It stores the image in the filesystem (`event/images`).
func UploadImage(api service.ImageAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body, that is the form data.
		// `10 << 20` specifies a maximum upload size of 10MB.
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			log.Printf("Failed to parse file from body: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Read the form data.
		file, fileHeader, err := r.FormFile(FORM_DATA_NAME)
		if err != nil {
			log.Printf("Failed to get form file: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Upload the image (save it in the `images` directory).
		u, err := api.UploadImage(file, fileHeader, r.Host)
		if err != nil {
			log.Printf("Failed to upload image: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Marshal the upload.
		bytes, err := json.Marshal(u)
		if err != nil {
			log.Printf("Failed to marshal upload result: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond JSON with the upload that has the url of our file.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}
