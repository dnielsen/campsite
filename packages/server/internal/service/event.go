package service

import (
	"log"
	"time"
)

type Event struct {
	ID            string         `json:"id,omitempty" gorm:"primaryKey;type:uuid"`
	Name          string         `json:"name,omitempty"`
	Description   string         `json:"description,omitempty"`
	StartDate     time.Time      `json:"startDate,omitempty"`
	EndDate       time.Time      `json:"endDate,omitempty"`
	Photo         string         `json:"photo,omitempty"`
	OrganizerName string         `json:"organizerName,omitempty"`
	Address       string         `json:"address,omitempty"`
	Sessions      []Session      `json:"sessions,omitempty"`
	Speakers      []Speaker      `json:"speakers,omitempty" gorm:"many2many:event_speakers;"`
}

func (api *api) GetEventById(id string) (*Event, error) {
	var event Event
	_ = api.db.Preload("Sessions").Preload("Speakers").Where("id = ?", id).First(&event)
	log.Println(event)
	return &event, nil
}