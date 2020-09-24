package service

import (
	"dave-web-app/packages/server/internal/config"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"gorm.io/gorm"
	"mime/multipart"
)

type API struct {
	db *gorm.DB
	u  *s3manager.Uploader
	c  *config.Config
}

// Create a new API object.
func NewAPI(db *gorm.DB, s *session.Session, c *config.Config) *API {
	u := s3manager.NewUploader(s)
	return &API{db, u, c}
}

type EventAPI interface {
	GetAllEvents() (*[]Event, error)
	CreateEvent(i EventInput) (*Event, error)
	GetEventById(id string) (*Event, error)
	DeleteEventById(id string) error
	EditEventById(id string, i EventInput) error
}

type SessionAPI interface {
	GetAllSessions() (*[]Session, error)
	CreateSession(i SessionInput) (*Session, error)
	GetSessionById(id string) (*Session, error)
	DeleteSessionById(id string) error
}

type SpeakerAPI interface {
	GetAllSpeakers() (*[]Speaker, error)
	CreateSpeaker(i SpeakerInput) (*Speaker, error)
	GetSpeakerById(id string) (*Speaker, error)
	DeleteSpeakerById(id string) error
}

type S3API interface {
	UploadImage(file multipart.File, fileHeader *multipart.FileHeader) (*Upload, error)
}
