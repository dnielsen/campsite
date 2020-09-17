package service

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Event struct {
	ID            string         `json:"id,omitempty" gorm:"primaryKey;type:uuid"`
	Name          string         `json:"name,omitempty"`
	Description   string         `json:"description,omitempty"`
	StartDate     *time.Time      `json:"startDate,omitempty"`
	EndDate       *time.Time      `json:"endDate,omitempty"`
	Photo         string         `json:"photo,omitempty"`
	OrganizerName string         `json:"organizerName,omitempty"`
	Address       string         `json:"address,omitempty"`
	Sessions      []Session      `json:"sessions,omitempty"`
	Speakers      []Speaker      `json:"speakers,omitempty" gorm:"many2many:event_speakers;"`
}

func (api *api) GetEventById(id string) (*Event, error) {
	var event Event
	res := api.db.Preload("Speakers").Preload("Sessions.Speakers").Where("id = ?", id).First(&event)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (api *api) GetAllEvents() (*[]Event, error) {
	var events []Event
	res := api.db.Find(&events)
	if err := res.Error; err != nil{
		return nil, err
	}
	return &events, nil
}


// Add a UUID automatically on creation so that we can skip it in our methods.
func (e *Event) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New().String()
	return
}

type EventInput struct {
	Name string `json:"name,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate *time.Time `json:"endDate,omitempty"`
	Description string `json:"description,omitempty"`
	Photo string `json:"photo,omitempty"`
	OrganizerName string `json:"organizerName,omitempty"`
	Address string `json:"address,omitempty"`
}

func (api *api) CreateEvent(i EventInput) (*Event, error) {
	// The ID will be added on insert.
	e := Event{
		Name:          i.Name,
		Description:   i.Description,
		StartDate:     i.StartDate,
		EndDate:       i.EndDate,
		Photo:         "https://apple.com",
		OrganizerName: "John Doe",
		Address:       "San Francisco, CA",
	}
	res := api.db.Create(&e)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func (api *api) EditEvent(id string, i EventInput) (*Event, error) {
	eventUpdates := &Event{
		Name:          i.Name,
		Description:   i.Description,
		StartDate:     i.StartDate,
		EndDate:       i.EndDate,
		Photo:         i.Photo,
		OrganizerName: i.OrganizerName,
		Address:       i.Address,
	}
	// Update the event in the database.
	if err := api.db.Model(&Event{}).Where("id = ?", id).Updates(&eventUpdates).Error; err != nil {
		return nil, err
	}

	// Grab the updated event from the database.
	var e Event
	if err := api.db.Where("id = ?", id).First(&e).Error; err != nil {
		return nil, err
	}
	return &e, nil
}




