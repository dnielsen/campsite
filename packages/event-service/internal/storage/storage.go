package storage

import (
	"dave-web-app/packages/event-service/internal/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
)

func NewS3Session(c *config.S3Config) *session.Session {
	s, err := session.NewSession(
		&aws.Config{
			Region: aws.String(c.Region),
		})
	if err != nil {
		log.Fatalf("Failed to create new aws session: %v", err)
	}
	log.Println("Connected to S3")

	return s
}
