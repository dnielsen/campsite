package service

import (
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/dnielsen/campsite/pkg/model"
	"gorm.io/gorm"
)

type API struct {
	db     *gorm.DB
	c      *config.Config
}

func NewAPI(db *gorm.DB, c *config.Config) *API {
	// We're using our custom `HttpClient` to enable mocking.
	return &API{db, c}
}

type SessionAPI interface {
	GetAllSessions() (*[]model.Session, error)
	CreateSession(i model.SessionInput) (*model.Session, error)
	GetSessionById(id string) (*model.Session, error)
	EditSessionById(id string, i model.SessionInput) (*model.Session, error)
	DeleteSessionById(id string) error
}
