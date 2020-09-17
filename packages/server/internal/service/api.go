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
	CreateEvent(i CreateEventInput) (*Event, error)
}

type SessionDatastore interface {
	GetSessionById(id string) (*Session, error)
	GetAllSessions() (*[]Session, error)
	CreateSession(i CreateSessionInput) (*Session, error)
}

type SpeakerDatastore interface {
	GetSpeakerById(id string) (*Speaker, error)
	GetAllSpeakers() (*[]Speaker, error)
	CreateSpeaker(i CreateSpeakerInput) (*Speaker, error)
}