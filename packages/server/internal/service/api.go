package service

import (
	"gorm.io/gorm"
)

type api struct {
	db *gorm.DB
}

func NewAPI(db *gorm.DB) *api {
	return &api{db}
}

type EventDatastore interface {
	GetEventById(id string) (*Event, error)
	GetAllEvents() (*[]Event, error)
	CreateEvent(i EventInput) (*Event, error)
	DeleteEvent(id string) error
}

type SessionDatastore interface {
	GetSessionById(id string) (*Session, error)
	GetAllSessions() (*[]Session, error)
	CreateSession(i SessionInput) (*Session, error)
	DeleteSession(id string) error
}

type SpeakerDatastore interface {
	GetSpeakerById(id string) (*Speaker, error)
	GetAllSpeakers() (*[]Speaker, error)
	CreateSpeaker(i SpeakerInput) (*Speaker, error)
	DeleteSpeaker(id string) error
}
