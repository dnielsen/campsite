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

type Datastore interface {
	GetEventById(id string) (*Event, error)
	GetAllSessions() (*[]Session, error)
	GetSessionsByIds(id []string) (*[]Session, error)
	GetSessionById(id string) (*Session, error)
	GetAllSpeakers() (*[]Speaker, error)
	GetSpeakersByIds(ids []string) (*[]Speaker, error)
	GetSpeakerById(id string) (*Speaker, error)
}
