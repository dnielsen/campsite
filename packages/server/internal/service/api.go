package service

import (
	"dave-web-app/packages/server/internal/config"
	"gorm.io/gorm"
	"mime/multipart"
	"os"
)

type API struct {
	db *gorm.DB
	c  *config.Config
}

// Create a new API object.
func NewAPI(db *gorm.DB, c *config.Config) *API {
	return &API{db, c}
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
	EditSessionById(id string, i SessionInput) error
}

type SpeakerAPI interface {
	GetAllSpeakers() (*[]Speaker, error)
	CreateSpeaker(i SpeakerInput) (*Speaker, error)
	GetSpeakerById(id string) (*Speaker, error)
	DeleteSpeakerById(id string) error
	EditSpeakerById(id string, i SpeakerInput) error
}

type ImageAPI interface {
	GetImage(filename string) (*os.File, error)
	UploadImage(file multipart.File, fileHeader *multipart.FileHeader, host string) (*Upload, error)
}
