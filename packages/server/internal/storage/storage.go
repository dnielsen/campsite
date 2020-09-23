package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
)

const (
	// TODO: move it to env and config
	AWS_S3_REGION = "eu-central-1"
)

func NewS3Session() *session.Session {
	s, err := session.NewSession(
		&aws.Config{
			Region: aws.String(AWS_S3_REGION),
		})
	if err != nil {
		log.Fatalf("Failed to create new aws session: %v", err)
	}
	log.Println("Connected to S3")

	return s
}
