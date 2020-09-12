package service

import (
	"gorm.io/gorm"
	"net/http"
)

type api struct {
	db *gorm.DB
	c HttpClient
}

func NewAPI(db *gorm.DB, c HttpClient) *api {
	return &api{db, c}
}

// We define our own interface so that we can mock it,
// and therefore test our fetch functions.
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Datastore interface {
	GetEventById(id string) (*Event, error)
	GetAllSessions() (*[]Session, error)
	GetSessionsByEventId(id string) (*[]Session, error)
	GetSessionById(id string) (*Session, error)
	GetAllSpeakers() (*[]Speaker, error)
	GetSpeakersByIds(ids []string) (*[]Speaker, error)
	GetSpeakerById(id string) (*Speaker, error)
}
