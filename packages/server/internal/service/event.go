package service

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Event struct {
	ID            string     `json:"id,omitempty" gorm:"primaryKey;type:uuid"`
	Name          string     `json:"name,omitempty"`
	Description   string     `json:"description,omitempty"`
	StartDate     *time.Time `json:"startDate,omitempty"`
	EndDate       *time.Time `json:"endDate,omitempty"`
	Photo         string     `json:"photo,omitempty"`
	OrganizerName string     `json:"organizerName,omitempty"`
	Address       string     `json:"address,omitempty"`
	Sessions      []Session  `json:"sessions"`
}

type EventInput struct {
	Name          string     `json:"name,omitempty" validate:"required,min=2,max=100"`
	// `gte` stands for >= time.Now.UTC()
	StartDate     *time.Time `json:"startDate,omitempty" validate:"required,gte"`
	EndDate       *time.Time `json:"endDate,omitempty" validate:"required,gte"`
	Description   string     `json:"description,omitempty" validate:"required,min=20,max=2000"`
	Photo         string     `json:"photo,omitempty" validate:"required,min=10,max=200"`
	OrganizerName string     `json:"organizerName,omitempty" validate:"required,min=2,max=50"`
	Address       string     `json:"address,omitempty" validate:"required,min=5,max=100"`
	// `min=1,max=100` == `1 <= len(SessionIds) >= 100`
	SessionIds []string `json:"sessionIds,omitempty" validate:"required,min=1,max=100"`
}

// Add a UUID automatically on creation so that we can skip it in our methods.
func (e *Event) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New().String()
	return
}

func (api *api) GetEventById(id string) (*Event, error) {
	var event Event
	res := api.db.Preload("Sessions.Speakers").Where("id = ?", id).First(&event)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (api *api) GetAllEvents() (*[]Event, error) {
	var events []Event
	res := api.db.Find(&events)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &events, nil
}


func (api *api) CreateEvent(i EventInput) (*Event, error) {
	// Get the sessions from the database to attach them to the event.
	var sessions []Session
	if err := api.db.Where("id IN ?", i.SessionIds).Find(&sessions).Error; err != nil {
		return nil, err
	}
	// Create the event with the sessions attached.
	event := Event{
		Name:          i.Name,
		Description:   i.Description,
		StartDate:     i.StartDate,
		EndDate:       i.EndDate,
		Photo:         i.Photo,
		OrganizerName: i.OrganizerName,
		Address:       i.Address,
		Sessions:      sessions,
	}
	if err := api.db.Create(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (api *api) EditEvent(id string, i EventInput) error {
	// Get the sessions from the database to attach them to the event.
	var sessions []Session
	if err := api.db.Where("id IN ?", i.SessionIds).Find(&sessions).Error; err != nil {
		return err
	}
	// Update the event in the database.
	event := &Event{
		ID: id,
		Name:          i.Name,
		Description:   i.Description,
		StartDate:     i.StartDate,
		EndDate:       i.EndDate,
		Photo:         i.Photo,
		OrganizerName: i.OrganizerName,
		Address:       i.Address,
		Sessions: sessions,
	}
	if err := api.db.Model(&event).Updates(&event).Error; err != nil {
		return err
	}
	return nil
}
