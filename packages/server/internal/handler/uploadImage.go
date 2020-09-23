package handler

import (
	"dave-web-app/packages/server/internal/service"
	"encoding/json"
	"log"
	"net/http"
)


func UploadImage(datastore service.S3Datastore) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		// `10 << 20` specifies a maximum upload of 10MB.
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			log.Printf("Failed to parse file from body: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			log.Printf("Failed to get form file: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		result, err := datastore.UploadImage(file, fileHeader)
		if err != nil {
			log.Printf("Failed to uploaded image: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		bytes, err := json.Marshal(result)
		if err != nil {
			log.Printf("Failed to marshal upload result: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}
