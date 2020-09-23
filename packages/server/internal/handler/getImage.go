package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

func GetImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the filename parameter.
		vars := mux.Vars(r)
		filename := vars[FILENAME]
		// Get the path to the image
		path := fmt.Sprintf("./images/%v", filename)

		// Open the image.
		img, err := os.Open(path)
		if err != nil {
			log.Printf("Failed to open image: %v", err)
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

