package service

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"gorm.io/gorm"
	"mime/multipart"
)

type api struct {
	db *gorm.DB
	u *s3manager.Uploader
}

func NewAPI(db *gorm.DB, s *session.Session) *api {
	u := s3manager.NewUploader(s)
	return &api{db, u}
}

type EventService interface {
	GetAllEvents() (*[]Event, error)
	CreateEvent(i EventInput) (*Event, error)
	GetEventById(id string) (*Event, error)
	DeleteEventById(id string) error
}

type SessionService interface {
	GetAllSessions() (*[]Session, error)
	CreateSession(i SessionInput) (*Session, error)
	GetSessionById(id string) (*Session, error)
	DeleteSessionById(id string) error
}

type SpeakerService interface {
	GetAllSpeakers() (*[]Speaker, error)
	CreateSpeaker(i SpeakerInput) (*Speaker, error)
	GetSpeakerById(id string) (*Speaker, error)
	DeleteSpeakerById(id string) error
}

type S3Service interface {
	UploadImage(file multipart.File, fileHeader *multipart.FileHeader) (*Upload, error)
}