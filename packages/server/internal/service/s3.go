package service

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"mime/multipart"
	)

// TODO: move it to env/config file.
const AWS_S3_BUCKET = "events-monolith"

type Upload struct {
	Url string `json:"url"`
}

func (api *api) UploadImage(file multipart.File, fileHeader *multipart.FileHeader) (*s3manager.UploadOutput, error) {
	buffer := make([]byte, fileHeader.Size)
	if _, err := file.Read(buffer); err != nil {
		return nil, err
	}

	result, err := api.u.Upload(&s3manager.UploadInput{
		Bucket:                    aws.String(AWS_S3_BUCKET),
		Key:                       aws.String(fileHeader.Filename),
		Body: 					   bytes.NewReader(buffer),
		ACL: aws.String("public-read"),
		ContentType: aws.String("image/jpeg"),
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

