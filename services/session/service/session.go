package service

import (
	"github.com/dnielsen/campsite/pkg/model"
	"github.com/google/uuid"
)

func (api *API) GetSessionById(id string) (*model.Session, error) {
	session := model.Session{ID: id}
	if err := api.db.Preload("Speakers").Preload("Event.Sessions").First(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (api *API) GetAllSessions() (*[]model.Session, error) {
	var sessions []model.Session
	if err := api.db.Find(&sessions).Error; err != nil {
		return nil, err
	}
	return &sessions, nil
}

func (api *API) DeleteSessionById(id string) error {
	if err := api.db.Where("id = ?", id).Delete(&model.Session{}).Error; err != nil {
		return err
	}
	return nil
}

func (api *API) CreateSession(i model.SessionInput) (*model.Session, error) {
	var speakers []model.Speaker
	if err := api.db.Where("id IN ?", i.SpeakerIds).Find(&speakers).Error; err != nil {
		return nil, err
	}
	session := model.Session{
		ID:          uuid.New().String(),
		Name:        i.Name,
		StartDate:   i.StartDate,
		EndDate:     i.EndDate,
		Description: i.Description,
		EventID:     i.EventId,
		Url:         i.Url,
		Speakers:    speakers,
	}
	if err := api.db.Create(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (api *API) EditSessionById(id string, i model.SessionInput) error {
	// Update the session (besides speakers).
	session := model.Session{
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
	var speakers []model.Speaker
	if err := api.db.Where("id IN ?", i.SpeakerIds).Select("id").Find(&speakers).Error; err != nil {
		return err
	}
	// Update the speakers
	if err := api.db.Model(&session).Association("Speakers").Replace(speakers); err != nil {
		return err
	}
	return nil
}
