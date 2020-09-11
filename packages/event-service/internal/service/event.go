package service

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Photo     string    `json:"photo"`
	// Temporarily there's no Organizer structure
	OrganizerName string `json:"organizerName"`
	Address       string `json:"address"`
}

type EventDatastore interface {
	GetEventById(id uuid.UUID) (*Event, error)
}

func (api *api) GetEventById(id uuid.UUID) (*Event, error) {
	var event Event
	_ = api.db.First(&event, id)
	return &event, nil
}
