package service

import (
	"time"
)

type Event struct {
	ID        string `gorm:"primaryKey;type:uuid" json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Photo     string    `json:"photo"`
	// Temporarily there's no Organizer structure
	OrganizerName string   `json:"organizerName"`
	Address       string   `json:"address"`
	SessionIds    []string `gorm:"type:uuid[]" json:"-"`
	SpeakerIds    []string `gorm:"type:uuid[]" json:"-"`
}

func (api *api) GetEventById(id uint) (*Event, error) {
	var event Event
	_ = api.db.First(&event, id)
	return &event, nil
}
