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
	// We define our own HttpClient to enable mocking (for easier testing).
	// We don't have tests yet, however, it's a common practice to do that
	// for this reason.
	return &API{db, c}
}

type EventAPI interface {
	GetAllEvents() (*[]model.Event, error)
	CreateEvent(i model.EventInput) (*model.Event, error)
	GetEventById(id string) (*model.Event, error)
	EditEventById(id string, i model.EventInput) (*model.Event, error)
	DeleteEventById(id string) error
}
