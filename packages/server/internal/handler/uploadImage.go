package handler

import (
	"dave-web-app/packages/server/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

// Our frontend appends a file and sets a `file` name.
// It's the most common used name.
const FORM_DATA_NAME = "file"

func UploadImage(datastore service.S3Service) http.HandlerFunc  {
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

		// Upload the image to S3.
		upload, err := datastore.UploadImage(file, fileHeader)
		if err != nil {
			log.Printf("Failed to uploaded image: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Marshal the upload.
		bytes, err := json.Marshal(upload)
		if err != nil {
			log.Printf("Failed to marshal upload result: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond JSON with the upload.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}