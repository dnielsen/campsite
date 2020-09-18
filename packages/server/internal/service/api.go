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
	EditEvent(id string, i EventInput) (*Event, error)
}

type SessionDatastore interface {
	GetSessionById(id string) (*Session, error)
	GetAllSessions() (*[]Session, error)
	CreateSession(i SessionInput) (*Session, error)
	EditSession(id string, i SessionInput) (*Session, error)
}

type SpeakerDatastore interface {
	GetSpeakerById(id string) (*Speaker, error)
	GetAllSpeakers() (*[]Speaker, error)
	CreateSpeaker(i SpeakerInput) (*Speaker, error)
	EditSpeaker(id string, i SpeakerInput) (*Speaker, error)
}
