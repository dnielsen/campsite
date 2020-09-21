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
	GetAllEvents() (*[]Event, error)
	CreateEvent(i EventInput) (*Event, error)
	GetEventById(id string) (*Event, error)
	DeleteEventById(id string) error
}

type SessionDatastore interface {
	GetAllSessions() (*[]Session, error)
	CreateSession(i SessionInput) (*Session, error)
	GetSessionById(id string) (*Session, error)
	DeleteSessionById(id string) error
}

type SpeakerDatastore interface {
	GetAllSpeakers() (*[]Speaker, error)
	CreateSpeaker(i SpeakerInput) (*Speaker, error)
	GetSpeakerById(id string) (*Speaker, error)
	DeleteSpeakerById(id string) error
}
