package service

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	ID          string    `gorm:"primaryKey;type:uuid" json:"id"`
	Name        string    `json:"name" gorm:"not null"`
	StartDate   *time.Time `json:"startDate" gorm:"not null"`
	EndDate     *time.Time `json:"endDate" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Url         string    `json:"url" gorm:"not null"`
	Event 		Event `json:"-"`
	EventID 	string `json:"-" gorm:"type:uuid;not null"`
	Speakers    []Speaker `json:"speakers,omitempty" gorm:"many2many:session_speakers;constraint:OnDelete:CASCADE;"`
}

type SessionInput struct {
	Name        string     `json:"name,omitempty"`
	StartDate   *time.Time `json:"startDate,omitempty"`
	EndDate     *time.Time `json:"endDate,omitempty"`
	Description string     `json:"description,omitempty"`
	Url         string     `json:"url,omitempty"`
	SpeakerIds []string `json:"speakerIds,omitempty"`
	EventId string `json:"eventId,omitempty"`
}

func (api *API) GetSessionById(id string) (*Session, error) {
	session := Session{ID: id}
	if err := api.db.Preload("Speakers").First(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (api *API) GetAllSessions() (*[]Session, error) {
	var sessions []Session
	if err := api.db.Find(&sessions).Error; err != nil {
		return nil, err
	}
	return &sessions, nil
}

func (api *API) DeleteSessionById(id string) error {
	if err := api.db.Where("id = ?", id).Delete(&Session{}).Error; err != nil {
		return err
	}
	return nil
}

func (api *API) CreateSession(i SessionInput) (*Session, error) {
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

func (api *API) EditSessionById(id string, i SessionInput) error {
	// Update the session (besides speakers).
	session := Session{
		ID:          id,
		Name:        i.Name,
		StartDate:   i.StartDate,
		EndDate:     i.EndDate,
		Description: i.Description,
		Url:         i.Url,
		EventID:     i.EventId,
	}
	if err := api.db.Updates(&session).Error; err != nil {
		return err
	}
	// Update the session speakers. We're doing it this way instead of adding it to the
	// session struct because otherwise we would just add new speaker ids to the session_speaker table
	// instead of replacing them.

	// Get the speakers with just their ids.
	var speakers []Speaker
	if err := api.db.Where("id IN ?", i.SpeakerIds).Select("id").Find(&speakers).Error; err != nil {
		return err
	}
	// Update the speakers
	if err := api.db.Model(&session).Association("Speakers").Replace(speakers); err != nil {
		return err
	}
	return nil
}