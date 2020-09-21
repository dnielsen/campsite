package service

import (
	"github.com/google/uuid"
	"time"
)
//
//type Session struct {
//	ID          string    `json:"id"`
//	Name        string    `json:"name"`
//	StartDate   *time.Time `json:"startDate"`
//	EndDate     *time.Time `json:"endDate"`
//	Description string    `json:"description"`
//	// Either live zoom or recorded video's youtube link.
//	Url string `json:"url"`
//	EventID 	string `json:"eventId"`
//	Event Event          `json:"event,omitempty"`
//	Speakers []Speaker `json:"speakers,omitempty"`
//}

type Session struct {
	ID          string    `gorm:"primaryKey;type:uuid" json:"id"`
	Name        string    `json:"name" gorm:"not null"`
	StartDate   *time.Time `json:"startDate" gorm:"not null"`
	EndDate     *time.Time `json:"endDate" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Url         string    `json:"url" gorm:"not null"`
	Event 		Event `json:"-"`
	EventID 	string `json:"-" gorm:"type:uuid;not null"`
	Speakers    []Speaker `json:"speakers,omitempty" gorm:"many2many:session_speakers;"`
}



type SessionInput struct {
	// Name is a required field with a minimum and maximum length of 2 and 100 respectively.
	Name        string     `json:"name,omitempty" validate:"required,min=2,max=100"`
	// `validate:"gte"` checks if the date is >= `time.Now.UTC()`
	StartDate   *time.Time `json:"startDate,omitempty" validate:"required,gte"`
	EndDate     *time.Time `json:"endDate,omitempty" validate:"required,gte"`
	Description string     `json:"description,omitempty" validate:"required,min=10,max=1000"`
	Url         string     `json:"url,omitempty" validate:"required,min=10,max=200"`
	// `validate:"min=1"` here means the length of the array must be >= 1.
	SpeakerIds []string `json:"speakerIds,omitempty" validate:"required,min=1,max=10"`
	EventId string `json:"eventId,omitempty" validate:"required,uuid4"`
}

func (api *api) GetSessionById(id string) (*Session, error) {
	session := Session{ID: id}
	if err := api.db.Preload("Speakers").First(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (api *api) GetAllSessions() (*[]Session, error) {
	var sessions []Session
	if err := api.db.Find(&sessions).Error; err != nil {
		return nil, err
	}
	return &sessions, nil
}

func (api *api) DeleteSessionById(id string) error {
	if err := api.db.Where("id = ?", id).Delete(&Session{}).Error; err != nil {
		return err
	}
	return nil
}

func (api *api) CreateSession(i SessionInput) (*Session, error) {
	var speakers []Speaker
	if err := api.db.Where("id IN ?", i.SpeakerIds).Find(&speakers).Error; err != nil {
		return nil, err
	}
	session := Session{
		ID:          uuid.New().String(),
		Name:        i.Name,
		StartDate:   i.StartDate,
		EndDate:     i.EndDate,
		Description: i.Description,
		EventID:     i.EventId,
		Url:         i.Url,
		Speakers: speakers,
	}
	if err := api.db.Create(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}