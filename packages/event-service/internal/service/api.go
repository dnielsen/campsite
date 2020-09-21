package service

import (
	"dave-web-app/packages/event-service/internal/config"
	"gorm.io/gorm"
	"net/http"
)

type api struct {
	db     *gorm.DB
	client HttpClient
	c      *config.Config
}

func NewAPI(db *gorm.DB, client HttpClient, c *config.Config) *api {
	return &api{db, client, c}
}

// We define our own interface so that we can mock it,
// and therefore test our fetch functions.
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type EventDatastore interface {
	GetEventById(id string) (*Event, error)
	GetAllEvents() (*[]Event, error)
	CreateEvent(i EventInput) (*Event, error)
	DeleteEventById(id string) error
}

type SessionDatastore interface {
	GetAllSessions() (*[]Session, error)
	GetSessionById(id string) (*Session, error)
	CreateSession(i SessionInput) (*Session, error)
	DeleteSessionById(id string) error
}

type SpeakerDatastore interface {
	GetAllSpeakers() (*[]Speaker, error)
	GetSpeakerById(id string) (*Speaker, error)
	CreateSpeaker(i SpeakerInput) (*Speaker, error)
	DeleteSpeakerById(id string) error
}