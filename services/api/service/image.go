package service

import (
	"fmt"
	"github.com/dnielsen/campsite/pkg/model"
	"io/ioutil"
	"mime/multipart"
	"os"
)

const IMAGES_DIRECTORY_PATH = "./images"

// If we were to make a production app, we'd store it, for example, using
// Amazon S3.
func (api *API) UploadImage(file multipart.File, fileHeader *multipart.FileHeader, host string) (*model.Upload, error) {
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

// Retrieves the image from the filesystem.
func (api *API) GetImage(filename string) (*os.File, error) {
	// Get the path to the image
	path := fmt.Sprintf("%v/%v", IMAGES_DIRECTORY_PATH, filename)

	// Open the image.
	img, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return img, nil
}
