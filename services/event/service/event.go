package service

import (
	"campsite/pkg/model"
	"github.com/google/uuid"
	"sort"
)

func (api *API) GetAllEvents() (*[]model.Event, error) {
	var events []model.Event
	if err := api.db.Order("start_date desc").Find(&events).Error; err != nil {
		return nil, err
	}
	return &events, nil
}

func (api *API) CreateEvent(i model.EventInput) (*model.Event, error) {
	e := model.Event{
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

func (api *API) GetEventById(id string) (*model.Event, error) {
	e := model.Event{ID: id}
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
	e := model.Event{ID: id}
	if err := api.db.Delete(&e).Error; err != nil {
		return err
	}
	return nil
}

func (api *API) EditEventById(id string, i model.EventInput) error {
	e := model.Event{
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
