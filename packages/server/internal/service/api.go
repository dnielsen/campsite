package service

import (
	"campsite/packages/server/internal/config"
	"gorm.io/gorm"
	"mime/multipart"
	"os"
)

type API struct {
	db     *gorm.DB
	c      *config.Config
}

func NewAPI(db *gorm.DB, c *config.Config) *API {
	return &API{db, c}
}

type EventAPI interface {
	GetAllEvents() (*[]Event, error)
	CreateEvent(i EventInput) (*Event, error)
	GetEventById(id string) (*Event, error)
	EditEventById(id string, i EventInput) error
	DeleteEventById(id string) error
}

type SessionAPI interface {
	GetAllSessions() (*[]Session, error)
	CreateSession(i SessionInput) (*Session, error)
	GetSessionById(id string) (*Session, error)
	EditSessionById(id string, i SessionInput) error
	DeleteSessionById(id string) error
}

type SpeakerAPI interface {
	GetAllSpeakers() (*[]Speaker, error)
	CreateSpeaker(i SpeakerInput) (*Speaker, error)
	GetSpeakerById(id string) (*Speaker, error)
	EditSpeakerById(id string, i SpeakerInput) error
	DeleteSpeakerById(id string) error
}

type ImageAPI interface {
	GetImage(filename string) (*os.File, error)
	UploadImage(file multipart.File, fileHeader *multipart.FileHeader, host string) (*Upload, error)
}