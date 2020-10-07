package service

import (
	"github.com/google/uuid"
	"sort"
	"time"
)

type Event struct {
	ID            string     `json:"id" gorm:"type:uuid"`
	Name          string     `json:"name" gorm:"not null"`
	Description   string     `json:"description" gorm:"not null"`
	RegistrationUrl string `json:"registrationUrl" gorm:"not null"`
	StartDate     *time.Time `json:"startDate" gorm:"not null"`
	EndDate       *time.Time `json:"endDate" gorm:"not null"`
	Photo         string     `json:"photo" gorm:"not null"`
	OrganizerName string     `json:"organizerName" gorm:"not null"`
	Address       *string    `json:"address"`
	Sessions      []Session  `json:"sessions"`
	Speakers []Speaker `json:"speakers,omitempty" gorm:"-"`
}

type EventInput struct {
	Name string `json:"name,omitempty"`
	StartDate     *time.Time `json:"startDate,omitempty"`
	EndDate       *time.Time `json:"endDate,omitempty"`
	RegistrationUrl string `json:"registrationUrl,omitempty"`
	Description   string     `json:"description,omitempty"`
	Photo         string     `json:"photo,omitempty"`
	OrganizerName string     `json:"organizerName,omitempty"`
	Address       *string    `json:"address,omitempty"`
}

func (api *API) GetAllEvents() (*[]Event, error) {
	var events []Event
	if err := api.db.Find(&events).Error; err != nil {
		return nil, err
	}
	return &events, nil
}

func (api *API) CreateEvent(i EventInput) (*Event, error) {
	e := Event{
		ID:            uuid.New().String(),
		Name:          i.Name,
		Description:   i.Description,
		StartDate:     i.StartDate,
		EndDate:       i.EndDate,
		Photo:         i.Photo,
		OrganizerName: i.OrganizerName,
		Address:       i.Address,
	}
	if err := api.db.Create(&e).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func (api *API) GetEventById(id string) (*Event, error) {
	e := Event{ID: id}
	// We're preloading sessions and sessions' speakers since we'll need them in the event by id page.
	// Right now we're getting all of the properties, we'll optimize the query later.
	if err := api.db.Preload("Sessions.Speakers").First(&e).Error; err != nil {
		return nil, err
	}

	// We're sorting the sessions by start date so that we don't need
	// to on the frontend. If our microservices were communicating with
	// each other, we could've handled that in the session service.
	// We're not sure if we can do that via a query in this case.
	sort.Slice(e.Sessions, func(i, j int) bool {
		return e.Sessions[i].StartDate.Before(*e.Sessions[j].StartDate)
	})

	return &e, nil
}

func (api *API) DeleteEventById(id string) error {
	e := Event{ID: id}
	if err := api.db.Delete(&e).Error; err != nil {
		return err
	}
	return nil
}

func (api *API) EditEventById(id string, i EventInput) error {
	e := Event{
		ID:            id,
		Name:          i.Name,
		Description:   i.Description,
		StartDate:     i.StartDate,
		EndDate:       i.EndDate,
		Photo:         i.Photo,
		OrganizerName: i.OrganizerName,
		Address:       i.Address,
	}
	if err := api.db.Updates(&e).Error; err != nil {
		return err
	}
	return nil
}