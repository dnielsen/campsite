package handler

import (
	"fmt"
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

// `/images/{filename}` GET route. It doesn't communicate with the database or any of the services.
// It retrieves the images from the filesystem (`event/images`).
func GetImageByFilename(c *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the filename parameter.
		vars := mux.Vars(r)
		filename := vars[FILENAME]
		// Get the image.
		img, err := findImage(filename)
		if err != nil {
			log.Printf("Failed to get image: %v", err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		defer img.Close()
		// Respond with the image.
		w.Header().Set(CONTENT_TYPE, IMAGE_JPEG)
		w.WriteHeader(http.StatusOK)
		io.Copy(w, img)
	}
}

// Retrieves the image from the filesystem.
func findImage(filename string) (*os.File, error) {
	// Get the path to the image
	path := fmt.Sprintf("%v/%v", IMAGES_DIRECTORY_PATH, filename)

	// Open the image.
	img, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return img, nil
}
