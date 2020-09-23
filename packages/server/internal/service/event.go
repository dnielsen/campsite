package service

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	ID            string     `json:"id" gorm:"type:uuid"`
	Name          string     `json:"name" gorm:"not null"`
	Description   string     `json:"description" gorm:"not null"`
	StartDate     *time.Time `json:"startDate" gorm:"not null"`
	EndDate       *time.Time `json:"endDate" gorm:"not null"`
	Photo         string     `json:"photo" gorm:"not null"`
	OrganizerName string     `json:"organizerName" gorm:"not null"`
	Address       *string    `json:"address"`
	Sessions      []Session  `json:"sessions"`
}

type EventInput struct {
	Name string `json:"name,omitempty" validate:"required,min=2,max=100"`
	// `gte` stands for >= time.Now.UTC()
	StartDate     *time.Time `json:"startDate,omitempty" validate:"required,gte"`
	EndDate       *time.Time `json:"endDate,omitempty" validate:"required,gte"`
	Description   string     `json:"description,omitempty" validate:"required,min=20,max=2000"`
	Photo         string     `json:"photo,omitempty" validate:"required,min=10,max=200"`
	OrganizerName string     `json:"organizerName,omitempty" validate:"required,min=2,max=50"`
	Address       *string    `json:"address,omitempty" validate:"required,min=5,max=100"`
}

func (api *API) GetAllEvents() (*[]Event, error) {
	var events []Event
	if err := api.db.Find(&events).Error; err != nil {
		return nil, err
	}
	return &events, nil
}

func (api *API) CreateEvent(i EventInput) (*Event, error) {
	event := Event{
		ID:            uuid.New().String(),
		Name:          i.Name,
		Description:   i.Description,
		StartDate:     i.StartDate,
		EndDate:       i.EndDate,
		Photo:         i.Photo,
		OrganizerName: i.OrganizerName,
		Address:       i.Address,
	}
	if err := api.db.Create(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (api *API) GetEventById(id string) (*Event, error) {
	event := Event{ID: id}
	// We're preloading sessions and sessions' speakers since we'll need them in the event by id page.
	// Right now we're getting all of the properties, we'll optimize the query later.
	if err := api.db.Preload("Sessions.Speakers").First(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (api *API) DeleteEventById(id string) error {
	event := Event{ID: id}
	if err := api.db.Delete(&event).Error; err != nil {
		return err
	}
	return nil
}
