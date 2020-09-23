package service

import (
	"dave-web-app/packages/event-service/internal/config"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"gorm.io/gorm"
	"mime/multipart"
	"net/http"
)

type API struct {
	db     *gorm.DB
	u *s3manager.Uploader
	client HttpClient
	c      *config.Config
}

func NewAPI(db *gorm.DB, s *session.Session, client HttpClient, c *config.Config) *API {
	u := s3manager.NewUploader(s)
	return &API{db, u, client, c}
}

// We define our own interface so that we can mock it,
// and therefore test our fetch functions.
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type EventAPI interface {
	GetEventById(id string) (*Event, error)
	GetAllEvents() (*[]Event, error)
	CreateEvent(i EventInput) (*Event, error)
	DeleteEventById(id string) error
}

type SessionAPI interface {
	GetAllSessions() (*[]Session, error)
	GetSessionById(id string) (*Session, error)
	CreateSession(i SessionInput) (*Session, error)
	DeleteSessionById(id string) error
}

type SpeakerAPI interface {
	GetAllSpeakers() (*[]Speaker, error)
	GetSpeakerById(id string) (*Speaker, error)
	CreateSpeaker(i SpeakerInput) (*Speaker, error)
	DeleteSpeakerById(id string) error
}

type S3API interface {
	UploadImage(file multipart.File, fileHeader *multipart.FileHeader) (*Upload, error)
}