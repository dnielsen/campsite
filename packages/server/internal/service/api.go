package service

import (
	"gorm.io/gorm"
)

type api struct {
	db     *gorm.DB
}

func NewAPI(db *gorm.DB) *api {
	return &api{db}
}

type EventDatastore interface {
	GetEventById(id string) (*Event, error)
	GetAllEvents() (*[]Event, error)
}

type SessionDatastore interface {
	GetSessionById(id string) (*Session, error)
	GetAllSessions() (*[]Session, error)
}

type SpeakerDatastore interface {
	GetSpeakerById(id string) (*Speaker, error)
	GetAllSpeakers() (*[]Speaker, error)
}