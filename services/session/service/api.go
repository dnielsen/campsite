package service

import (
	"campsite/pkg/config"
	"campsite/pkg/model"
	"gorm.io/gorm"
	"net/http"
)

type API struct {
	db     *gorm.DB
	client HttpClient
	c      *config.Config
}

func NewAPI(db *gorm.DB, c *config.Config) *API {
	// We're using our custom `HttpClient` to enable mocking.
	var client HttpClient = http.DefaultClient

	return &API{db, client, c}
}

// We define our own interface so that we can mock it,
// and therefore test our fetch functions.
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type SessionAPI interface {
	GetAllSessions() (*[]model.Session, error)
	CreateSession(i model.SessionInput) (*model.Session, error)
	GetSessionById(id string) (*model.Session, error)
	EditSessionById(id string, i model.SessionInput) error
	DeleteSessionById(id string) error
}
