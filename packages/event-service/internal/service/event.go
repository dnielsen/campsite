package service

import (
	"github.com/lib/pq"
	"time"
)

type Event struct {
	ID            string         `gorm:"primaryKey;type:uuid" json:"id"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	StartDate     time.Time      `json:"startDate"`
	EndDate       time.Time      `json:"endDate"`
	Photo         string         `json:"photo"`
	OrganizerName string         `json:"organizerName"`
	Address       string         `json:"address"`
	SessionIds    pq.StringArray `json:"sessionIds" gorm:"type:uuid[]"`
	Sessions      []Session      `json:"sessions,omitempty" gorm:"-"`
	SpeakerIds    pq.StringArray `json:"speakerIds" gorm:"type:uuid[]" `
	Speakers      []Speaker      `json:"speakers,omitempty" gorm:"-"`
}

func (api *api) GetEventById(id string) (*Event, error) {
	var event Event
	_ = api.db.Where("id = ?", id).First(&event)
	return &event, nil
}
