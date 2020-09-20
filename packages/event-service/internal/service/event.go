package service

import (
	"time"
)

type Event struct {
	ID            string         `gorm:"primaryKey;type:uuid" json:"id"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	StartDate     *time.Time      `json:"startDate"`
	EndDate       *time.Time      `json:"endDate"`
	Photo         string         `json:"photo"`
	OrganizerName string         `json:"organizerName"`
	Address       string         `json:"address"`
	Sessions      []Session      `json:"sessions,omitempty" gorm:"-"`
	Speakers      []Speaker      `json:"speakers,omitempty" gorm:"-"`
}

func (api *api) GetEventById(id string) (*Event, error) {
	var event Event
	if err := api.db.Where("id = ?", id).First(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (api *api) GetAllEvents() (*[]Event, error) {
	var events []Event
	if err := api.db.Find(&events).Error; err != nil {
		return nil, err
	}
	return &events, nil
}