package service

import (
	"dave-web-app/packages/session-service/internal/config"
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

type SessionDatastore interface {
	GetAllSessions() (*[]Session, error)
	CreateSession(i SessionInput) (*Session, error)
	GetSessionById(id string) (*Session, error)
	DeleteSessionById(id string) error
}
