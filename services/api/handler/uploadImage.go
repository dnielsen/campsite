package handler

import (
	"encoding/json"
	"fmt"
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/dnielsen/campsite/pkg/jwt"
	"github.com/dnielsen/campsite/pkg/model"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

// Our frontend appends a file and sets the form data name to `file`.
// It's the most commonly used name for form data.
const FORM_DATA_NAME = "file"

// `/images` POST route. It doesn't communicate with the database or any of the services.
// It stores the image in the filesystem (`event/images`).
func UploadImage(c *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Verify the JWT token since it's a protected route.
		tokenCookie, err := r.Cookie(c.Jwt.CookieName)
		if err != nil {
			log.Printf("Failed to get cookie: %v", err)
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		_, err = jwt.VerifyToken(tokenCookie.Value, &c.Jwt)
		if err != nil {
			log.Printf("Failed to verify token: %v", err)
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
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
		u, err := storeImage(file, fileHeader, r.Host)
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
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}

// If we were to make a production app, we'd store it, for example, using Amazon S3.
func storeImage(file multipart.File, fileHeader *multipart.FileHeader, host string) (*model.Upload, error) {
	originalFilename := fileHeader.Filename
	// Create a temporary file in our `images` directory with a unique filename
	// so that we can later save the received file into it. `*` will be replaced
	// with a random string, so that our saved image file has a unique name.
	// Example: `man.jpg` => `345834858-man.jpg`
	tempFile, err := ioutil.TempFile(IMAGES_DIRECTORY_PATH, fmt.Sprintf("*-%v", originalFilename))
	if err != nil {
		return nil, err
	}
	defer tempFile.Close()

	// Read the received file.
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Write the file bytes into our temporary file.
	if _, err := tempFile.Write(bytes); err != nil {
		return nil, err
	}

	// Return with an Upload
	path := tempFile.Name()
	u := model.Upload{Url: fmt.Sprintf("http://%v/%v", host, path)}
	return &u, nil
}