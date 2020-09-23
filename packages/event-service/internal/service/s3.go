package service

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/globalsign/mgo/bson"
	"mime/multipart"
	"path/filepath"
)

const (
	AWS_CONTENT_TYPE = "image/*"     // Accept only images.
	AWS_ACL          = "public-read" // Let anyone access (view) the image.
)

// Get the properties we want by defining our struct.
// Also, we can set json encoding, that let's us make
// the properties camelCase. Otherwise the output would be, say:
// `{"Location": "https://amazon.com", ...}`
type Upload struct {
	Url string `json:"url"`
}

func (api *API) UploadImage(file multipart.File, fileHeader *multipart.FileHeader) (*Upload, error) {
	// Create a unique filename for the file with a proper extension.
	filename := fmt.Sprintf("images/%v%v", bson.NewObjectId().Hex(), filepath.Ext(fileHeader.Filename))

	// Read the file data and put it in a buffer.
	buffer := make([]byte, fileHeader.Size)
	if _, err := file.Read(buffer); err != nil {
		return nil, err
	}

	// Upload the file to Amazon S3.
	result, err := api.u.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(api.c.S3.Bucket),
		Key:         aws.String(filename),
		Body:        bytes.NewReader(buffer),
		ACL:         aws.String(AWS_ACL),
		ContentType: aws.String(AWS_CONTENT_TYPE),
	})
	if err != nil {
		return nil, err
	}

	// Return the url to the file.
	upload := Upload{Url: result.Location}
	return &upload, nil
}