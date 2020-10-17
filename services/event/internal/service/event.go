package service

import (
	"github.com/google/uuid"
)


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
	return &e, nil
}

func (api *API) DeleteEventById(id string) error {
	e := Event{ID: id}
	if err := api.db.Delete(&e).Error; err != nil {
		return err
	}
	return nil
}

func (api *API) EditEventById(id string, i EventInput) (*Event, error) {
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
		return nil, err
	}
	return &e, nil
}
